package auth

import (
	"context"
	"time"

	"github.com/dizzrt/dauth/internal/domain/auth/cache"
	"github.com/dizzrt/dauth/internal/domain/auth/entity"
	"github.com/dizzrt/dauth/internal/infra/foundation"
)

var _ cache.AuthorizationCodeCache = (*AuthorizationCodeCacheImpl)(nil)

const _PREFIX = "auth:code:"

type AuthorizationCodeCacheImpl struct {
	*foundation.RedisClient
}

func NewAuthorizationCodeCacheImpl(cli *foundation.RedisClient) cache.AuthorizationCodeCache {
	return &AuthorizationCodeCacheImpl{
		RedisClient: cli,
	}
}

func (impl *AuthorizationCodeCacheImpl) Key(code string) string {
	return _PREFIX + code
}

func (impl *AuthorizationCodeCacheImpl) Set(ctx context.Context, code string, value *entity.AuthorizationCode, ttl time.Duration) error {
	return impl.Cmdable().Set(ctx, impl.Key(code), value, ttl).Err()
}

func (impl *AuthorizationCodeCacheImpl) Get(ctx context.Context, code string) (*entity.AuthorizationCode, error) {
	var value entity.AuthorizationCode
	if err := impl.Cmdable().Get(ctx, impl.Key(code)).Scan(&value); err != nil {
		return nil, err
	}

	return &value, nil
}

func (impl *AuthorizationCodeCacheImpl) Delete(ctx context.Context, code string) error {
	return impl.Cmdable().Del(ctx, impl.Key(code)).Err()
}
