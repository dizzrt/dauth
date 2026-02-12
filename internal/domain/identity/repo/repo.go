package repo

import (
	"context"

	"github.com/dizzrt/dauth/internal/domain/identity/entity"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *entity.User) (uint32, error)
	GetUserByID(ctx context.Context, uid uint32) (*entity.User, error)
}
