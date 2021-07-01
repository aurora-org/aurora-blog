package server

import (
	"aurora/blog/api/app/service/manager/internal/service"
	"github.com/gin-gonic/gin"
)

// data -> biz -> service -> server

// NewHttpRouter set gin.Engine as http.Handler
func NewHttpRouter(service *service.Service) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// set router
	httpRouter(router, service)
	return router
}

func httpRouter(app *gin.Engine, service *service.Service) {
	v := app.Group("/v1")
	// hello router
	{
		v.GET("/hello", service.Hello)
	}
}
