package data

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"aurora/blog/api/app/service/manager/internal/data/vo"
	"entgo.io/ent/examples/start/ent"

	//where "aurora/blog/api/app/service/manager/internal/data/ent/user"
	"context"
	"go.uber.org/zap"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *zap.SugaredLogger
}

// NewHelloRepo ...
func NewUserRepo(data *Data, logger *zap.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  logger.Sugar(),
	}
}

func (r *userRepo) Create(ctx context.Context, vo *vo.UserVO) (*ent.User, error) {
	u, err := r.data.db.User.
		Create().
		SetName(vo.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}
