package service

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"github.com/google/wire"
)

// ProvideSet for service package ...
var ProvideSet = wire.NewSet(NewService)

func NewService(biz *biz.Biz) *Service {
	return &Service{
		biz: biz,
	}
}
