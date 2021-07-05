package data

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"go.uber.org/zap"
)

var _ biz.ArticleRepo = (*articleRepo)(nil)

type articleRepo struct {
	data *Data
	log  *zap.SugaredLogger
}

// NewArticleRepo ...
func NewArticleRepo(data *Data, logger *zap.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  logger.Sugar(),
	}
}
