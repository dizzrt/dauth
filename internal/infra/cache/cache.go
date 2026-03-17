package cache

import (
	"github.com/dizzrt/dauth/internal/infra/cache/impl"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	impl.NewTokenCacheImpl,
)
