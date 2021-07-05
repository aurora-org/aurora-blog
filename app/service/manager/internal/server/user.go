package server

import (
	"aurora/blog/api/app/service/manager/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(group *gin.RouterGroup, service *service.UserService) {

	group.POST("/users", service.CreateUser)
	group.GET("/users", service.GetAllUser)
	group.GET("/users/:id", service.GetUserByID)

}
