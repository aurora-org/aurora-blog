package biz

type Repo interface {
	Hello() (string, error)
}

var _ Repo = (*Biz)(nil)

type Biz struct {
	repo Repo
}

func (b Biz) Hello() (string, error) {
	return b.repo.Hello()
}
