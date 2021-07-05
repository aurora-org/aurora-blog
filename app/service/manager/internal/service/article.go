package service

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"go.uber.org/zap"
)

type ArticleService struct {
	biz *biz.ArticleBusiness
	log *zap.SugaredLogger
}

func NewArticleService(biz *biz.ArticleBusiness, logger *zap.Logger) *ArticleService {
	return &ArticleService{
		biz: biz,
		log: logger.Sugar(),
	}
}