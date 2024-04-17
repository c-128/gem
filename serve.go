package gem

import (
	"crypto/tls"
	"errors"
	"net"
)

func ServeAndListen(address string, handler Handler, certFile string, keyFile string) error {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	config := &tls.Config{
		ClientAuth: tls.RequestClientCert,
		MinVersion: tls.VersionTLS12,
		Certificates: []tls.Certificate{
			cert,
		},
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		tlsConn := tls.Server(conn, config)
		go handleConnection(tlsConn, handler)
	}
}

func handleConnection(conn *tls.Conn, handler Handler) error {
	defer conn.Close()

	err := conn.Handshake()
	if err != nil {
		return err
	}

	var cert *Cert = nil
	connState := conn.ConnectionState()
	connCerts := connState.PeerCertificates

	if len(connCerts) == 1 {
		cert = newCertFromX509(connCerts[0])
	}

	header := make([]byte, maxURILength+2)
	read, err := conn.Read(header)
	if err != nil {
		return err
	}

	crLfIndex := -1
	for i := 0; i < read-1; i++ {
		if header[i] == '\r' && header[i+1] == '\n' {
			crLfIndex = i
			break
		}
	}

	if crLfIndex == -1 {
		return errors.New("failed to parse url")
	}

	rawURL := string(header[:crLfIndex])
	url, err := ParseURL(rawURL)
	if err != nil {
		return err
	}

	ctx := &Ctx{
		url:    url,
		cert:   cert,
		conn:   conn,
		params: make(map[string]string),

		status:      nil,
		info:        nil,
		wroteStatus: false,

		locals: make(map[string]any),
	}

	handler(ctx)

	return nil
}
