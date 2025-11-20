package foundation

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewBaseDB,
	NewLogger,
	NewRegistrar,
	NewTracerProvider,
)
