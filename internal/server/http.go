package server

import (
	"github.com/dizzrt/dauth/api/gen/user"
	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/dauth/internal/handler"
	"github.com/dizzrt/ellie/log"
	"github.com/dizzrt/ellie/middleware/tracing"
	"github.com/dizzrt/ellie/transport/http"
)

func NewHTTPServer(c *conf.Bootstrap, logger log.LogWriter, userHandler *handler.UserHandler) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			tracing.TracingMiddleware(),
		),
	}

	httpServerConf := c.Server.HTTP

	if httpServerConf.Addr != "" {
		opts = append(opts, http.Address(httpServerConf.Addr))
	}

	srv := http.NewServer(opts...)
	user.RegisterUserServiceHTTPServer(srv, userHandler)

	return srv
}
