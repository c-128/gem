package main

import (
	"fmt"

	"github.com/c-128/gem"
)

// Simple static route
func handleIndex(ctx *gem.Ctx) error {
	ctx.Status(gem.StatusSuccess, "text/gemini")

	fmt.Fprintf(ctx, "# Hello gemini!\n")

	return nil
}

// Simple static route
func handleRoute(ctx *gem.Ctx) error {
	ctx.Status(gem.StatusSuccess, "text/gemini")

	fmt.Fprintf(ctx, "# Hello route!\n")

	return nil
}

// Simple dynamic route
func handleDynamic(ctx *gem.Ctx) error {
	ctx.Status(gem.StatusSuccess, "text/gemini")
	dynamic := ctx.Param("dynamic")

	fmt.Fprintf(ctx, "# Hello gemini!\nYou entered: %s", dynamic)

	return nil
}

func main() {
	mux := gem.NewMux()

	// Register routes
	mux.Handle("/", handleIndex)
	mux.Handle("/route", handleRoute)
	mux.Handle("/dynamic/:dynamic", handleDynamic)

	// Listen at the address and serve the handler
	err := gem.ServeAndListen(
		"0.0.0.0:1965",
		mux.Handler,
		"cert/cert.crt",
		"cert/key.key",
	)
	if err != nil {
		panic(err)
	}
}
