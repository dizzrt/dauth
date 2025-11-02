//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/dizzrt/dauth/internal/application"
	"github.com/dizzrt/dauth/internal/conf"
	user_biz "github.com/dizzrt/dauth/internal/domain/user/biz"
	"github.com/dizzrt/dauth/internal/handler"
	"github.com/dizzrt/dauth/internal/infra/repo"
	"github.com/dizzrt/dauth/internal/server"
	"github.com/dizzrt/ellie"
	"github.com/dizzrt/ellie/log"
	"github.com/google/wire"
)

func wireApp(bootstrap *conf.Bootstrap, logger log.LogWriter) (*ellie.App, func(), error) {
	panic(wire.Build(newApp, server.ProviderSet, handler.ProviderSet, application.ProviderSet, user_biz.ProviderSet, repo.ProviderSet))
}
