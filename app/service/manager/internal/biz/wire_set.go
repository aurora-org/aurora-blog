package biz

import (
	"github.com/google/wire"
)

// ProvideSet for biz package ...
var ProvideSet = wire.NewSet(NewUserBusiness)

func NewUserBusiness(repo UserRepo) *UserBusiness {
	return &UserBusiness{repo: repo}
}
