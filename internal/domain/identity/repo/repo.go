package repo

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/domain/identity/entity"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, uid uint32) (*entity.User, error)
	ListUsers(ctx context.Context, page, size int32) ([]*entity.User, int64, error)
	UpdateUserStatus(ctx context.Context, uid uint32, status identity.UserStatus) error
}
