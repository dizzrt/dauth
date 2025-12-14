package token

import (
	"context"
	"errors"
	"time"

	"github.com/dizzrt/dauth/internal/domain/token/cache"
	"github.com/dizzrt/dauth/internal/infra/foundation"
	"github.com/dizzrt/ellie/log"
	"github.com/redis/go-redis/v9"
)

var _ cache.TokenRevokeCache = (*TokenRevokeCacheImpl)(nil)

const _PREFIX = "token:revoked:"

type TokenRevokeCacheImpl struct {
	*foundation.RedisClient
}

func NewTokenRevokeCacheImpl(cli *foundation.RedisClient) cache.TokenRevokeCache {
	return &TokenRevokeCacheImpl{
		RedisClient: cli,
	}
}

func (impl *TokenRevokeCacheImpl) Key(token string) string {
	return _PREFIX + token
}

func (impl *TokenRevokeCacheImpl) Revoke(ctx context.Context, token string, reason string, expiresAt time.Time) error {
	ttl := time.Until(expiresAt)
	return impl.Cmdable().Set(ctx, impl.Key(token), reason, ttl).Err()
}

func (impl *TokenRevokeCacheImpl) IsRevoked(ctx context.Context, token string) (bool, string, error) {
	v := impl.Cmdable().Get(ctx, impl.Key(token))

	err := v.Err()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, "", nil
		}

		log.CtxErrorf(ctx, "get token revoke cache failed: %s", err.Error())

		// when get cache failed, suppose token has been revoked
		return true, err.Error(), err
	}

	return true, v.Val(), nil
}
