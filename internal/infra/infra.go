package infra

import (
	"github.com/dizzrt/dauth/internal/infra/foundation"
	"github.com/dizzrt/dauth/internal/infra/repo"
	"github.com/dizzrt/dauth/internal/infra/utils/security/jwt"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	foundation.ProviderSet,
	repo.ProviderSet,
	jwt.NewJWTManager,
)
