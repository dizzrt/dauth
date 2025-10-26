package server

import (
	"github.com/dizzrt/dauth/api/gen/example"
	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/dauth/internal/handler"
	"github.com/dizzrt/ellie/log"
	"github.com/dizzrt/ellie/transport/http"
)

func NewHTTPServer(c *conf.Bootstrap, logger log.LogWriter, exampleHandler *handler.ExampleHandler) *http.Server {
	opts := []http.ServerOption{}

	httpServerConf := c.Server.HTTP

	if httpServerConf.Addr != "" {
		opts = append(opts, http.Address(httpServerConf.Addr))
	}

	srv := http.NewServer(opts...)
	example.RegisterExampleHTTPServer(srv, exampleHandler)

	return srv
}
