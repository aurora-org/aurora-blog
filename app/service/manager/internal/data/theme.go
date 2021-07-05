package data

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"go.uber.org/zap"
)

var _ biz.ThemeRepo = (*themeRepo)(nil)

type themeRepo struct {
	data *Data
	log  *zap.SugaredLogger
}

// NewThemeRepo ...
func NewThemeRepo(data *Data, logger *zap.Logger) biz.ThemeRepo {
	return &themeRepo{
		data: data,
		log:  logger.Sugar(),
	}
}
