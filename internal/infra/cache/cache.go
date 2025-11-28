package cache

import (
	"github.com/dizzrt/dauth/internal/infra/cache/impl/auth"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	auth.NewAuthorizationCodeCacheImpl,
)
