package biz

import (
	"context"

	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"github.com/dizzrt/dauth/internal/domain/identity/repo"
)

var _ UserBiz = (*userBiz)(nil)

type UserBiz interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, uid uint32) (*entity.User, error)
}

type userBiz struct {
	userRepo repo.UserRepo
}

func NewUserBiz(userRepo repo.UserRepo) UserBiz {
	return &userBiz{
		userRepo: userRepo,
	}
}

func (biz *userBiz) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (biz *userBiz) GetUser(ctx context.Context, uid uint32) (*entity.User, error) {
	return nil, nil
}
