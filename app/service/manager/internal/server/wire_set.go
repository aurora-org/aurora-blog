package server

import (
	"aurora/blog/api/pkg/lifecycle"
	"aurora/blog/api/pkg/registry"
	"github.com/google/wire"
)

// ProvideSet for server package ...
var ProvideSet = wire.NewSet(NewHttpServer, registry.New, NewHttpRouter)

func NewHttpServer() lifecycle.Server {
	return &HttpServer{}
}
