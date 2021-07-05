package service

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"aurora/blog/api/app/service/manager/internal/data/vo"
	"aurora/blog/api/pkg/functions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type UserService struct {
	biz *biz.UserBusiness
	log *zap.SugaredLogger
}

func NewUserService(biz *biz.UserBusiness, logger *zap.Logger) *UserService {
	return &UserService{
		biz: biz,
		log: logger.Sugar(),
	}
}

func (u *UserService) CreateUser(ctx *gin.Context) {
	user := &vo.UserVO{}

	err := ctx.BindJSON(user)
	saved, err := u.biz.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": saved,
	})
}

func (u *UserService) GetUserByID(ctx *gin.Context) {
	id := functions.StringToInt(ctx.Param("id"))
	if id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
		})
	}

	user, err := u.biz.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (u *UserService) GetAllUser(ctx *gin.Context) {
	user, err := u.biz.GetAllUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users": user,
	})
}
