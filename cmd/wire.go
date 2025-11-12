//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/dizzrt/dauth/internal/application"
	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/dauth/internal/domain/identity"
	"github.com/dizzrt/dauth/internal/handler"
	"github.com/dizzrt/dauth/internal/infra"
	"github.com/dizzrt/dauth/internal/infra/repo"
	"github.com/dizzrt/dauth/internal/server"
	"github.com/dizzrt/ellie"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/sdk/trace"
)

type WireApp struct {
	App *ellie.App
	TP  *trace.TracerProvider
}

func newWireApp(app *ellie.App, tp *trace.TracerProvider) *WireApp {
	return &WireApp{
		App: app,
		TP:  tp,
	}
}

func wireApp() (*WireApp, func(), error) {
	panic(wire.Build(
		newWireApp,
		newApp,
		conf.ProviderSet,
		server.ProviderSet,
		handler.ProviderSet,
		application.ProviderSet,
		identity.ProviderSet,
		repo.ProviderSet,
		infra.ProviderSet,
	))
}
