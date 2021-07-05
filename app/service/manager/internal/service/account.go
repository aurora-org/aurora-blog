package service

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"go.uber.org/zap"
)

type AccountService struct {
	biz *biz.AccountBusiness
	log *zap.SugaredLogger
}

func NewAccountService(biz *biz.AccountBusiness, logger *zap.Logger) *AccountService {
	return &AccountService{
		biz: biz,
		log: logger.Sugar(),
	}
}
