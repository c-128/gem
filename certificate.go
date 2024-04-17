package gem

import (
	"crypto/sha512"
	"crypto/x509"
	"encoding/hex"
	"time"
)

func newCertFromX509(cert *x509.Certificate) *Cert {
	now := time.Now()
	valid := true

	if now.After(cert.NotAfter) {
		valid = false
	}

	if now.Before(cert.NotBefore) {
		valid = false
	}

	return &Cert{
		valid:   valid,
		rawCert: cert,
	}
}

// A client certificate that is sent to the server.
type Cert struct {
	valid   bool
	rawCert *x509.Certificate
}

// Returns true if the certificate is within the valid time range.
func (cert *Cert) Valid() bool {
	return cert.valid
}

// Calculates the SHA512 sum of the certificate and encodes it in hex.
func (cert *Cert) Fingerprint() string {
	sum := sha512.Sum512(cert.rawCert.Raw)
	return hex.EncodeToString(sum[:])
}

// Returns the underlying x509 certificate.
func (cert *Cert) RawCert() *x509.Certificate {
	return cert.rawCert
}
