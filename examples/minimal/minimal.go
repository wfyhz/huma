package main

import (
	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/cli"
	"github.com/danielgtaylor/huma/responses"
)

func main() {
	app := cli.NewRouter("Minimal Example", "1.0.0")

	app.Resource("/").Get("get-root", "Get a short text message",
		responses.OK().ContentType("text/plain"),
	).Run(func(ctx huma.Context) {
		ctx.Header().Set("Content-Type", "text/plain")
		ctx.Write([]byte("Hello, world"))
	})

	app.Run()
}
