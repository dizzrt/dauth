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
	if ac.Tracing.Endpoint == "" {
		return nil, func() {}, nil
	}

	tp, err := tracing.Initialize(
		context.Background(),
		tracing.ServiceName(conf.Service),
		tracing.ServiceVersion(conf.Version),
		tracing.Endpoint(ac.Tracing.Endpoint),
		tracing.EndpointType(tracing.ParseEndpointType(ac.Tracing.EndpointType)),
		tracing.Insecure(ac.Tracing.Insecure),
		tracing.Metadata(map[string]string{
			"ip":          ac.Address,
			"env":         ac.ENV,
			"hostname":    conf.Hostname,
			"instance_id": conf.ServiceID,
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
