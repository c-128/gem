package gem

import (
	"crypto/tls"
	"errors"
	"fmt"
)

type Ctx struct {
	url    *URL
	cert   *Cert
	conn   *tls.Conn
	params map[string]string

	status      *int
	info        *string
	wroteStatus bool

	locals map[string]any
}

// Get the request URL.
func (ctx *Ctx) URL() URL {
	return *ctx.url
}

// Get the URL query.
func (ctx *Ctx) Query() string {
	return ctx.url.Query
}

// Get the client certificate.
func (ctx *Ctx) Cert() *Cert {
	return ctx.cert
}

// Get the underlying connection.
func (ctx *Ctx) Connection() *tls.Conn {
	return ctx.conn
}

// Get a route parameter.
func (ctx *Ctx) Param(key string) string {
	return ctx.params[key]
}

// Lookup a route parameter.
func (ctx *Ctx) LookupParam(key string) (string, bool) {
	param, found := ctx.params[key]
	return param, found
}

// Sets the status that will be send upon the first .Write call.
func (ctx *Ctx) Status(status int, info string) {
	ctx.status = &status
	ctx.info = &info
}

// Sends the status to the client.
// If the status has already been sent, it will error.
// If the status isn't between [10; 99], it will error.
func (ctx *Ctx) SendStatus(status int, info string) error {
	if ctx.wroteStatus {
		return errors.New("wrote status already")
	}

	if 10 > status || status > 99 {
		return errors.New("status must be between [10; 99]")
	}

	_, err := fmt.Fprintf(ctx.conn, "%d %s\r\n", status, info)
	if err != nil {
		return err
	}

	ctx.wroteStatus = true
	return nil
}

// Get a local.
// Will return nil if the local hasn't been found.
func (ctx *Ctx) Local(key string) any {
	local, found := ctx.locals[key]
	if !found {
		return nil
	}

	return local
}

// Set a local.
func (ctx *Ctx) SetLocal(key string, value any) {
	ctx.locals[key] = value
}

// Write a repsonse to the client.
// If the response status hasn't been sent yet using .SendStatus, this will send it.
// By default it responds with status code 20 and info "text/plain".
// These values can be overriden by calling .Status.
func (ctx *Ctx) Write(buffer []byte) (int, error) {
	if !ctx.wroteStatus {
		status := StatusSuccess
		info := "text/plain"

		if ctx.status != nil {
			status = *ctx.status
		}

		if ctx.info != nil {
			info = *ctx.info
		}

		err := ctx.SendStatus(status, info)
		if err != nil {
			return 0, err
		}
	}

	wrote, err := ctx.conn.Write(buffer)
	if err != nil {
		return wrote, err
	}

	return wrote, nil
}
