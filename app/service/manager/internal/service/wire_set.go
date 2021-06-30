package service

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"github.com/google/wire"
)

// ProvideSet for service package ...
var ProvideSet = wire.NewSet(NewService)

func NewService(repo biz.Repo) *Service {
	return &Service{
		repo: repo,
	}
}
