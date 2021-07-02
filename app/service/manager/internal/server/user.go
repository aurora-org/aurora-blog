package server

import (
	"aurora/blog/api/app/service/manager/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(group *gin.RouterGroup, service *service.UserService) {
	v := group.Group("/user")

	v.GET("/", service.CreateUser)
}

