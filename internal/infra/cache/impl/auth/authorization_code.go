package auth

import (
	"context"
	"time"

	"github.com/dizzrt/dauth/internal/domain/auth/cache"
	"github.com/dizzrt/dauth/internal/domain/auth/entity"
	"github.com/dizzrt/dauth/internal/infra/foundation"
)

var _ cache.AuthorizationCodeCache = (*AuthorizationCodeCacheImpl)(nil)

type AuthorizationCodeCacheImpl struct {
	cli *foundation.RedisClient
}

func NewAuthorizationCodeCacheImpl(cli *foundation.RedisClient) cache.AuthorizationCodeCache {
	return &AuthorizationCodeCacheImpl{
		cli: cli,
	}
}

func (impl *AuthorizationCodeCacheImpl) Set(ctx context.Context, code string, value *entity.AuthorizationCode, ttl time.Duration) error {
	return impl.cli.Cmdable().Set(ctx, code, value, ttl).Err()
}

func (impl *AuthorizationCodeCacheImpl) Get(ctx context.Context, code string) (*entity.AuthorizationCode, error) {
	var value entity.AuthorizationCode
	if err := impl.cli.Cmdable().Get(ctx, code).Scan(&value); err != nil {
		return nil, err
	}

	return &value, nil
}

func (impl *AuthorizationCodeCacheImpl) Delete(ctx context.Context, code string) error {
	return impl.cli.Cmdable().Del(ctx, code).Err()
}
