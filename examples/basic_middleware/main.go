package main

import (
	"github.com/c-128/gem"
)

// Basic handler to return a 51 error
func helloGemini(ctx *gem.Ctx) error {
	return gem.NotFoundErr
}

// Respond to the client with the code if it is a gem error
func errorHandler(handler gem.Handler) gem.Handler {
	return func(ctx *gem.Ctx) error {
		toHandle := handler(ctx)

		status := gem.StatusTemporaryFailure
		message := toHandle.Error()

		// If the error is a gem.Error, send another status
		switch err := toHandle.(type) {
		case gem.Error:
			status = err.Status()
		}

		ctx.SendStatus(status, message)
		return nil
	}
}

func main() {
	err := gem.ServeAndListen(
		"0.0.0.0:1965",
		// Wrap the handler with the middleware
		errorHandler(helloGemini),
		"cert/cert.crt",
		"cert/key.key",
	)
	if err != nil {
		panic(err)
	}
}
