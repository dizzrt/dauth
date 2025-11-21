package biz

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"github.com/dizzrt/dauth/internal/domain/identity/repo"
)

var _ UserBiz = (*userBiz)(nil)

type UserBiz interface {
	Authenticate(ctx context.Context, account string, password string) (*entity.User, error)
	UpdateLastLoginTime(ctx context.Context, uid uint32) error
	CreateUser(ctx context.Context, user *entity.User) (uint32, error)
	GetUserByID(ctx context.Context, uid uint32) (*entity.User, error)
	UpdateUserStatus(ctx context.Context, uid uint32, status identity.User_Status) error
	UpdateUserPassword(ctx context.Context, uid uint32, newPassword string) error
}

type userBiz struct {
	userRepo repo.UserRepo
}

func NewUserBiz(userRepo repo.UserRepo) UserBiz {
	return &userBiz{
		userRepo: userRepo,
	}
}

// only support Authenticate by email yet
func (biz *userBiz) Authenticate(ctx context.Context, account string, password string) (*entity.User, error) {
	user, err := biz.userRepo.GetUserByEmail(ctx, account)
	if err != nil {
		return nil, err
	}

	if err := user.VerifyPassword(password); err != nil {
		return nil, err
	}

	return user, nil
}

func (biz *userBiz) UpdateLastLoginTime(ctx context.Context, uid uint32) error {
	// TODO: update last login time
	return nil
}

func (biz *userBiz) CreateUser(ctx context.Context, user *entity.User) (uint32, error) {
	rpwd := user.Password
	pwd, err := GeneratePasswordHash(rpwd)
	if err != nil {
		return 0, err
	}

	user.Password = string(pwd)
	return biz.userRepo.CreateUser(ctx, user)
}

func (biz *userBiz) GetUserByID(ctx context.Context, uid uint32) (*entity.User, error) {
	return biz.userRepo.GetUserByID(ctx, uid)
}

func (biz *userBiz) UpdateUserStatus(ctx context.Context, uid uint32, status identity.User_Status) error {
	return biz.userRepo.UpdateUserStatus(ctx, uid, status)
}

func (biz *userBiz) UpdateUserPassword(ctx context.Context, uid uint32, password string) error {
	pwd, err := GeneratePasswordHash(password)
	if err != nil {
		return err
	}

	return biz.userRepo.UpdateUserPassword(ctx, uid, string(pwd))
}
