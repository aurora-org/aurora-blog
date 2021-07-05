package service

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"go.uber.org/zap"
)

type ThemeService struct {
	biz *biz.ThemeBusiness
	log *zap.SugaredLogger
}

func NewThemeService(biz *biz.ThemeBusiness, logger *zap.Logger) *ThemeService {
	return &ThemeService{
		biz: biz,
		log: logger.Sugar(),
	}
}
