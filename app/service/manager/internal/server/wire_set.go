package server

import (
	"aurora/blog/api/pkg/transport"
	"aurora/blog/api/pkg/transport/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// ProvideSet for server package ...
var ProvideSet = wire.NewSet(NewHttpServer, NewHttpRouter)

func NewHttpServer(engine *gin.Engine) transport.Server {
	return http.NewServer(engine)
}
