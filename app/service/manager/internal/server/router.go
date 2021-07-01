package server

import (
	"aurora/blog/api/app/service/manager/internal/service"
	"aurora/blog/api/pkg/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// call chain: data -> biz -> service -> server
// NewHttpRouter set gin.Engine as http.Handler
func NewHttpRouter(service *service.Service, logger *zap.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	engine.Use(middleware.Ginzap(logger, ""))
	engine.Use(middleware.RecoveryWithZap(logger, true))
	// set router
	httpRouter(engine, service)
	return engine
}

func httpRouter(engine *gin.Engine, service *service.Service) {
	v := engine.Group("/v1")
	// hello router
	{
		v.GET("/hello", service.Hello)
	}
}
