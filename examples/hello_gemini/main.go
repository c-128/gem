package main

import (
	"fmt"

	"github.com/c-128/gem"
)

// Basic handler
func helloGemini(ctx *gem.Ctx) error {
	// Set content type and status
	ctx.Status(gem.StatusSuccess, "text/gemini")

	// Send a simple "Hello gemini" to the client
	// This will also send a status set with .Status
	// If it isn't set, it will default to status code 20 and "text/plain"
	fmt.Fprintf(ctx, "# Hello gemini!\n")

	return nil
}

func main() {
	// Listen at the address and serve the handler
	err := gem.ServeAndListen(
		"0.0.0.0:1965",
		helloGemini,
		"cert/cert.crt",
		"cert/key.key",
	)
	if err != nil {
		panic(err)
	}
}
