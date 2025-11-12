package infra

import (
	"github.com/dizzrt/dauth/internal/infra/common"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	common.NewLogger,
	common.NewBaseDB,
	common.NewTracerProvider,
)
