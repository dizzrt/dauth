package foundation

import (
	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/ellie/contrib/registry/consul"
	"github.com/dizzrt/ellie/registry"
	"github.com/hashicorp/consul/api"
)

func NewRegistrar(ac *conf.AppConfig) registry.Registrar {
	if ac.Registry.Addr == "" {
		return nil
	}

	cli, err := api.NewClient(&api.Config{
		Address: ac.Registry.Addr,
	})

	if err != nil {
		panic(err)
	}

	return consul.New(cli)
}
