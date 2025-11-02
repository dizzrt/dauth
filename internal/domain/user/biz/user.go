package biz

import (
	"context"

	"github.com/dizzrt/dauth/internal/domain/user/entity"
	"github.com/dizzrt/dauth/internal/domain/user/repo"
	"github.com/dizzrt/ellie/log"
	"golang.org/x/crypto/bcrypt"
)

var _ UserBiz = (*userBiz)(nil)

type UserBiz interface {
	// create user, return uid when success
	CreateUser(ctx context.Context, user *entity.User) (uint, error)
	GetUserByID(ctx context.Context, uid uint32) (*entity.User, error)
}

type userBiz struct {
	userRepo repo.UserRepo
}

func NewUserBiz(userRepo repo.UserRepo) UserBiz {
	return &userBiz{
		userRepo: userRepo,
	}
}

func (biz *userBiz) CreateUser(ctx context.Context, user *entity.User) (uint, error) {
	log.CtxInfof(ctx, "[CreateUser] biz req: %v", user)

	pwd := user.Password
	pwdBytes := []byte(pwd)
	hashBytes, err := bcrypt.GenerateFromPassword(pwdBytes, 10)
	if err != nil {
		return 0, err
	}

	// save hash password
	user.Password = string(hashBytes)
	uid, err := biz.userRepo.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}

	return uid, nil
}

func (biz *userBiz) GetUserByID(ctx context.Context, uid uint32) (*entity.User, error) {
	log.CtxInfof(ctx, "[GetUserByID] biz req: %v", uid)

	user, err := biz.userRepo.GetUserByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return user, nil
}
