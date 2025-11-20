package foundation

import (
	"context"
	"time"

	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/ellie/middleware/tracing"
	trace_sdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func NewTracerProvider(ac *conf.AppConfig) (trace.TracerProvider, func(), error) {
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

	cleanup := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if sdkTp, ok := tp.(*trace_sdk.TracerProvider); ok {
			sdkTp.Shutdown(ctx)
		}
	}

	return tp, cleanup, nil
}
