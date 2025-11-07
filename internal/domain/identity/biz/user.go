package biz

import (
	"github.com/dizzrt/dauth/internal/domain/identity/repo"
)

var _ UserBiz = (*userBiz)(nil)

type UserBiz interface {
	// only support email yet
	// Authenticate(ctx context.Context, account string, password string) (*entity.User, error)
	// CreateUser(ctx context.Context, user *entity.User) (uint32, error)
	// GetUserByID(ctx context.Context, uid uint32) (*entity.User, error)
	// UpdateUserStatus(ctx context.Context, uid uint32, status identity.UserStatus) error
	// UpdateUserPassword(ctx context.Context, uid uint32, oldPassword, newPassword string) error
}

type userBiz struct {
	userRepo repo.UserRepo
}

func NewUserBiz(userRepo repo.UserRepo) UserBiz {
	return &userBiz{
		userRepo: userRepo,
	}
}
