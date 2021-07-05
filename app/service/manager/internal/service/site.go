package service

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"go.uber.org/zap"
)

type SiteService struct {
	biz *biz.SiteBusiness
	log *zap.SugaredLogger
}

func NewSiteService(biz *biz.SiteBusiness, logger *zap.Logger) *SiteService {
	return &SiteService{
		biz: biz,
		log: logger.Sugar(),
	}
}