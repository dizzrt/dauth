package repo

import (
	"context"
	"time"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/domain/identity/entity"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *entity.User) (uint32, error)
	GetUserByID(ctx context.Context, uid uint32) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUserPassword(ctx context.Context, uid uint32, password string) error
	UpdateUserStatus(ctx context.Context, uid uint32, status identity.User_Status) error
	UpdateLastLoginTime(ctx context.Context, uid uint32, lastLoginTime time.Time) error
}

type RoleRepo interface {
}

type UserRoleAssociationRepo interface {
}
