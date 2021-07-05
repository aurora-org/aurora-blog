package service

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"go.uber.org/zap"
)

type TagService struct {
	biz *biz.TagBusiness
	log *zap.SugaredLogger
}

func NewTagService(biz *biz.TagBusiness, logger *zap.Logger) *TagService {
	return &TagService{
		biz: biz,
		log: logger.Sugar(),
	}
}