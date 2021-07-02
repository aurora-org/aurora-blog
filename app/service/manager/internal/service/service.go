package service

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloService struct {
	biz *biz.Biz
}

func NewHelloService(biz *biz.Biz) *HelloService {
	return &HelloService{
		biz: biz,
	}
}


func (s *HelloService) Hello(ctx *gin.Context) {
	str, err := s.biz.HelloWorld()
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, str)
}
