package server

import (
	"aurora/blog/api/pkg/lifecycle"
	"aurora/blog/api/pkg/registry"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// ProvideSet for server package ...
var ProvideSet = wire.NewSet(NewHttpServer, NewRegistrar, NewHttpRouter)

func NewHttpServer(engine *gin.Engine) lifecycle.Server {
	return &HttpServer{
		engine: engine,
	}
}

func NewRegistrar() lifecycle.Registrar {
	return registry.New()
}