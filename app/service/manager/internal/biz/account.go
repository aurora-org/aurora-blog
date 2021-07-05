package biz

type AccountRepo interface {
}

type AccountBusiness struct {
	repo AccountRepo
}
