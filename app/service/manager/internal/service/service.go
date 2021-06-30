package service

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Service struct {
	biz *biz.Biz
}

func (s *Service) Hello(ctx *gin.Context) {
	str, err := s.biz.Hello()
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, str)
}
