package common

import (
	"context"

	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/ellie/middleware/tracing"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewTracerProvider(bootstrap *conf.Bootstrap) *trace.TracerProvider {
	// TODO read from conf
	tp, err := tracing.Initialize(
		context.Background(),
		tracing.ServiceName("dauth"),
		tracing.ServiceVersion("dev"),
		tracing.Endpoint("infra.dauth.com:4317"),
		tracing.EndpointType(tracing.EndpointType_GRPC),
		tracing.Insecure(true),
		tracing.Metadata(map[string]string{
			"ip":  "127.0.0.1",
			"env": "dev",
		}),
	)

	if err != nil {
		panic(err)
	}

	return tp
}
