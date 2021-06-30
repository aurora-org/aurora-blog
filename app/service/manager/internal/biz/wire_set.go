package biz

import (
	"aurora/blog/api/app/service/manager/internal/data"
	"github.com/google/wire"
)

// ProvideSet for biz package ...
var ProvideSet = wire.NewSet(NewBiz)

func NewBiz(data data.Data) Repo {
	return &Biz{data: data}
}
