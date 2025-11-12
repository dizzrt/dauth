package common

import (
	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/ellie/contrib/registry/consul"
	"github.com/dizzrt/ellie/registry"
	"github.com/hashicorp/consul/api"
)

func NewRegistrar(bootstrap *conf.Bootstrap) registry.Registrar {
	if bootstrap.Registry.Addr == "" {
		return nil
	}

	cli, err := api.NewClient(&api.Config{
		Address: bootstrap.Registry.Addr,
	})

	if err != nil {
		panic(err)
	}

	return consul.New(cli)
}
