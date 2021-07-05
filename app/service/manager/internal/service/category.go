package service

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"go.uber.org/zap"
)

type CategoryService struct {
	biz *biz.CategoryBusiness
	log *zap.SugaredLogger
}

func NewCategoryService(biz *biz.CategoryBusiness, logger *zap.Logger) *CategoryService {
	return &CategoryService{
		biz: biz,
		log: logger.Sugar(),
	}
}