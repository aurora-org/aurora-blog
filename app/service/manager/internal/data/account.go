package data

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"go.uber.org/zap"
)

var _ biz.AccountRepo = (*accountRepo)(nil)

type accountRepo struct {
	data *Data
	log  *zap.SugaredLogger
}

// NewAccountRepo ...
func NewAccountRepo(data *Data, logger *zap.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  logger.Sugar(),
	}
}
