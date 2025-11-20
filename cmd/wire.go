//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/dizzrt/dauth/internal/application"
	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/dauth/internal/domain"
	"github.com/dizzrt/dauth/internal/handler"
	"github.com/dizzrt/dauth/internal/infra"
	"github.com/dizzrt/dauth/internal/server"
	"github.com/dizzrt/ellie"
	"github.com/google/wire"
)

func wireApp() (*ellie.App, func(), error) {
	panic(wire.Build(
		conf.ProviderSet,
		infra.ProviderSet,
		domain.ProviderSet,
		application.ProviderSet,
		handler.ProviderSet,
		server.ProviderSet,
		newApp,
	))
}
