package data

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"go.uber.org/zap"
)

var _ biz.TagRepo = (*tagRepo)(nil)

type tagRepo struct {
	data *Data
	log  *zap.SugaredLogger
}

// NewTagRepo ...
func NewTagRepo(data *Data, logger *zap.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  logger.Sugar(),
	}
}