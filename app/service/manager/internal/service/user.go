package service

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"aurora/blog/api/app/service/manager/internal/data/vo"
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
	name := ctx.Query("name")
	u.log.Info(name)
	user := &vo.UserVO{
		Name: name,
	}
	saved := u.biz.CreateUser(ctx, user)
	ctx.JSON(http.StatusOK, gin.H{
		"user": saved,
	})
}
