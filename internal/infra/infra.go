package infra

import (
	"github.com/dizzrt/dauth/internal/infra/common"
	"github.com/dizzrt/dauth/internal/infra/repo"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	repo.ProviderSet,
	common.NewLogger,
	common.NewBaseDB,
	common.NewTracerProvider,
	common.NewRegistrar,
)
