package cache

import (
	"context"
	"time"

	"github.com/dizzrt/dauth/internal/domain/auth/entity"
)

type AuthorizationCodeCache interface {
	Set(ctx context.Context, code string, value *entity.AuthorizationCode, ttl time.Duration) error
	Get(ctx context.Context, code string) (*entity.AuthorizationCode, error)
	Delete(ctx context.Context, code string) error
}
