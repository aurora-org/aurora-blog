package biz

type Repo interface {
	Hello() (string, error)
}

type Biz struct {
	repo Repo
}

func (b Biz) HelloWorld() (string, error) {
	return b.repo.Hello()
}
