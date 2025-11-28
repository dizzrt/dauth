package foundation

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewBaseDB,
	NewRedisClient,
	NewLogger,
	NewRegistrar,
	NewTracerProvider,
)
