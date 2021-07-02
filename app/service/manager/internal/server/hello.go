package server

import (
	"aurora/blog/api/app/service/manager/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterHelloRouter(group *gin.RouterGroup, service *service.HelloService) {
	v := group.Group("/hello")

	v.GET("/", service.Hello)
}
