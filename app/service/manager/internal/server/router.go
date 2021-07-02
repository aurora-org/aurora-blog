package server

import (
	"aurora/blog/api/app/service/manager/internal/service"
	"aurora/blog/api/pkg/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// call chain: data -> biz -> service -> server
// NewHttpRouter set gin.Engine as http.Handler
func NewHttpRouter(service *service.UserService, logger *zap.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	engine.Use(middleware.Ginzap(logger))
	engine.Use(middleware.RecoveryWithZap(logger, true))
	// set router
	mng := engine.Group("/v1/mng")
	RegisterUserRouter(mng, service)
	return engine
}
