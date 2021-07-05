package data

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"go.uber.org/zap"
)

var _ biz.SiteRepo = (*siteRepo)(nil)

type siteRepo struct {
	data *Data
	log  *zap.SugaredLogger
}

// NewSiteRepo ...
func NewSiteRepo(data *Data, logger *zap.Logger) biz.SiteRepo {
	return &siteRepo{
		data: data,
		log:  logger.Sugar(),
	}
}