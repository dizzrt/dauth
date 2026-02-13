package repo

import (
	"github.com/dizzrt/dauth/internal/infra/repo/core"
	"github.com/dizzrt/dauth/internal/infra/repo/impl/identity"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	core.NewRepoCore,
	identity.NewUserRepoImpl,
)
