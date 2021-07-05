package data

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"go.uber.org/zap"
)

var _ biz.CategoryRepo = (*categoryRepo)(nil)

type categoryRepo struct {
	data *Data
	log  *zap.SugaredLogger
}

// NewCategoryRepo ...
func NewCategoryRepo(data *Data, logger *zap.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  logger.Sugar(),
	}
}