package biz

import (
	"aurora/blog/api/app/service/manager/internal/data/ent"
	"aurora/blog/api/app/service/manager/internal/data/vo"
)

type UserRepo interface {
	Create(vo *vo.UserVO) (*ent.User, error)
	GetByID(id int) (*ent.User, error)
	GetAll() ([]*ent.User, error)
}

type UserBusiness struct {
	repo UserRepo
}

func (b *UserBusiness) CreateUser(vo *vo.UserVO) (*ent.User, error) {
	user, err := b.repo.Create(vo)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (b *UserBusiness) GetUserByID(id int) (*ent.User, error) {
	user, err := b.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (b *UserBusiness) GetAllUser() ([]*ent.User, error) {
	user, err := b.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return user, nil
}
