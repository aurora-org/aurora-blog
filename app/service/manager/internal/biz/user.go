package biz

import (
	"aurora/blog/api/app/service/manager/internal/data/vo"
	"context"
	"entgo.io/ent/examples/start/ent"
)

type UserRepo interface {
	Create(ctx context.Context, vo *vo.UserVO) (*ent.User, error)
}

type UserBusiness struct {
	repo UserRepo
}

func (b *UserBusiness) CreateUser(ctx context.Context, vo *vo.UserVO) *ent.User {
	user, err := b.repo.Create(ctx, vo)
	if err != nil {
		return nil
	}
	return user
}
