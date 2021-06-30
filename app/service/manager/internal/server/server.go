package server

import (
	"aurora/blog/api/app/service/manager/internal/service"
	"github.com/gin-gonic/gin"
)

// data -> biz -> service -> server

// NewHttpRouter set gin.Engine as http.Handler
func NewHttpRouter(service service.Service) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// set book router
	bookRouter(router, service)
	return router
}

func bookRouter(app *gin.Engine, service service.Service) {
	app.GET("/hello", service.Hello)
}
