package biz

import "aurora/blog/api/app/service/manager/internal/data"

type Repo interface {
	Hello() (string, error)
}

var _ Repo = (*Biz)(nil)

type Biz struct {
	data data.Data
}

func (b Biz) Hello() (string, error) {
	return b.data.Hello()
}
