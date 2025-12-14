package cache

import (
	"github.com/dizzrt/dauth/internal/infra/cache/impl/auth"
	"github.com/dizzrt/dauth/internal/infra/cache/impl/token"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	token.NewTokenRevokeCacheImpl,
	auth.NewAuthorizationCodeCacheImpl,
)
