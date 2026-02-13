package biz

import (
	"context"
	"time"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/domain/identity/dto"
	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"github.com/dizzrt/dauth/internal/domain/identity/repo"
	"github.com/dizzrt/dauth/internal/infra/utils/security"
	"github.com/dizzrt/ellie/log"
)

var _ UserBiz = (*userBiz)(nil)

type UserBiz interface {
	CreateUser(ctx context.Context, user *dto.CreateUserDTO) (*entity.User, error)
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

func (biz *userBiz) CreateUser(ctx context.Context, user *dto.CreateUserDTO) (*entity.User, error) {
	pwd, err := security.GeneratePasswordHash(user.Password)
	if err != nil {
		log.CtxErrorf(ctx, "[Identity] generate password hash failed, input password: %s, err: %v", user.Password, err)
		return nil, err
	}

	now := time.Now()
	entityUser := &entity.User{
		Password:  pwd,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Phone:     user.Phone,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Status:    identity.UserStatus_ENABLED,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err = biz.userRepo.CreateUser(ctx, entityUser); err != nil {
		log.CtxErrorf(ctx, "[Identity] create user failed, err: %v", err)
		return nil, err
	}

	return entityUser, nil
}

func (biz *userBiz) GetUser(ctx context.Context, uid uint32) (*entity.User, error) {
	return nil, nil
}
