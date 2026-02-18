package persistence

import (
	"github.com/dizzrt/dauth/internal/infra/persistence/core"
	"github.com/dizzrt/dauth/internal/infra/persistence/impl/identity"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	core.NewRepoCore,
	identity.NewUserRepoImpl,
)
