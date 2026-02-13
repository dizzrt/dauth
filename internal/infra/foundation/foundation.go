package foundation

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewDB,
	NewRedisClient,
	NewLogger,
	NewRegistrar,
	NewTracerProvider,
)
