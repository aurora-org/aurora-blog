package main

import (
	"aurora/blog/api/pkg/lifecycle"
)

func newApp(server lifecycle.Server, registrar lifecycle.Registrar) *lifecycle.App {
	return lifecycle.New(
		lifecycle.WithName("manager"),
		lifecycle.WithVersion("v1.0"),
		lifecycle.WithMetadata(map[string]string{}),
		lifecycle.WithServer(server),
		lifecycle.WithRegistrar(registrar))
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
