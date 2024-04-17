# 💎 Gem
A library to quickly build gemini capsules using go.

## Installation
Installing Gem can be done easily using `go get`:
```sh
go get github.com/c-128/gem
```

## Examples
```go
package main

import (
	"io"

	"github.com/c-128/gem"
)

func helloGemini(ctx *gem.Ctx) error {
	ctx.Status(gem.StatusSuccess, "text/gemini")

	_, err := io.WriteString(ctx, "# Hello gemini!\n")
	if err != nil {
		return err
	}

	return nil
}

func main() {
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

```
Check out more examples in the [`examples`](examples) folder.
