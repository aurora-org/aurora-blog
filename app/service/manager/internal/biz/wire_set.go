package biz

import (
	"github.com/google/wire"
)

// ProvideSet for biz package ...
var ProvideSet = wire.NewSet(NewBiz)

func NewBiz(repo Repo) *Biz {
	return &Biz{repo: repo}
}
