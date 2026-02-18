package infra

import (
	"github.com/dizzrt/dauth/internal/infra/cache"
	"github.com/dizzrt/dauth/internal/infra/foundation"
	"github.com/dizzrt/dauth/internal/infra/persistence"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	foundation.ProviderSet,
	persistence.ProviderSet,
	cache.ProviderSet,
	// jwt.NewJWTManager,
)
