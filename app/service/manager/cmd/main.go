package main

import (
	"aurora/blog/api/pkg/lifecycle"
	"aurora/blog/api/pkg/transport"
	"go.uber.org/zap"
	"net/url"
)

func newApp(server transport.Server, logger *zap.Logger) *lifecycle.App {
	return lifecycle.New(
		lifecycle.WithName("manager"),
		lifecycle.WithVersion("1.0"),
		lifecycle.WithEndpoint(&url.URL{
			Host:        "127.0.0.1",
		}),
		lifecycle.WithMetadata(map[string]string{}),
		lifecycle.WithLogger(logger),
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
