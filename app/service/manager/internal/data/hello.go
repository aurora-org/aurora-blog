package data

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"go.uber.org/zap"
)

var _ biz.Repo = (*HelloRepo)(nil)

type HelloRepo struct {
	data *Data
	log  *zap.SugaredLogger
}

// NewHelloRepo ...
func NewHelloRepo(data *Data, logger *zap.Logger) biz.Repo {
	return &HelloRepo{
		data: data,
		log:  logger.Sugar(),
	}
}

func (h *HelloRepo) Hello() (string, error) {
	// call db
	return "hello world", nil
}
