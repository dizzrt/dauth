package impl

import (
	"context"
	"errors"
	"time"

	"github.com/dizzrt/dauth/internal/infra/foundation"
	"github.com/dizzrt/dauth/internal/infra/utils/cache"
	"github.com/dizzrt/ellie/log"
	"github.com/redis/go-redis/v9"
)

var _ cache.TokenCache = (*tokenCacheImpl)(nil)

const _PREFIX = "token:"

type tokenCacheImpl struct {
	*foundation.RedisClient
}

func NewTokenCacheImpl(cli *foundation.RedisClient) cache.TokenCache {
	return &tokenCacheImpl{
		RedisClient: cli,
	}
}

func (impl *tokenCacheImpl) key(token string) string {
	return _PREFIX + token
}

func (impl *tokenCacheImpl) Revoke(ctx context.Context, token string, reason string, expiresAt time.Time) error {
	ttl := time.Until(expiresAt)
	return impl.Cmdable().Set(ctx, impl.key(token), reason, ttl).Err()
}

func (impl *tokenCacheImpl) IsRevoked(ctx context.Context, token string) (bool, string, error) {
	v := impl.Cmdable().Get(ctx, impl.key(token))

	err := v.Err()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, "", nil
		}

		log.CtxErrorf(ctx, "get token revoke cache failed: %s", err.Error())

		// when get token failed, suppose token has been revoked
		return true, err.Error(), err
	}

	// cache hit, token has been revoked
	return true, v.Val(), nil
}
