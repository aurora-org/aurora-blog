package data

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"aurora/blog/api/app/service/manager/internal/data/ent"
	"aurora/blog/api/app/service/manager/internal/data/vo"
	"aurora/blog/api/pkg/constant"
	"time"

	where "aurora/blog/api/app/service/manager/internal/data/ent/user"
	"context"
	"go.uber.org/zap"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *zap.SugaredLogger
}

// NewUserRepo ...
func NewUserRepo(data *Data, logger *zap.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  logger.Sugar(),
	}
}

func (r *userRepo) Create(vo *vo.UserVO) (*ent.User, error) {
	user := &ent.User{
		Avater:   vo.Avater,
		Email:    vo.Email,
		Sex:      constant.SexMap[vo.Sex],
		Birthday: time.Now(),
		Name:     vo.Name,
		Nickname: vo.Nickname,
	}
	return r.data.db.User.Create().
		SetSex(user.Sex).
		SetName(user.Name).
		SetAvater(user.Avater).
		SetBirthday(user.Birthday).
		SetEmail(user.Email).
		SetNickname(user.Nickname).
		SetExtra(user.Extra).
		Save(context.Background())
}

func (r *userRepo) GetByID(id int) (*ent.User, error) {
	return r.data.db.User.Query().Where(where.IDEQ(id)).First(context.Background())
}

func (r *userRepo) GetAll() ([]*ent.User, error) {
	return r.data.db.User.Query().All(context.Background())
}
