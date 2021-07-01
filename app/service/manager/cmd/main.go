package main

import (
	"aurora/blog/api/pkg/lifecycle"
	"aurora/blog/api/pkg/transport"
)

func newApp(server transport.Server) *lifecycle.App {
	return lifecycle.New(
		lifecycle.WithName("manager"),
		lifecycle.WithVersion("v1.0"),
		lifecycle.WithMetadata(map[string]string{}),
		lifecycle.WithServer(server))
}

func main() {
	app, cleanup, err := initApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
