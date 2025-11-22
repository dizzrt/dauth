package server

import (
	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/dauth/internal/handler"
	"github.com/dizzrt/ellie/log"
	"github.com/dizzrt/ellie/middleware/tracing"
	"github.com/dizzrt/ellie/transport/http"
)

func NewHTTPServer(c *conf.AppConfig, logger log.LogWriter, identityHandler *handler.IdentityHandler, tokenHandler *handler.TokenHandler, clientHandler *handler.ClientHandler) *http.Server {
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
	identity.RegisterUserServiceHTTPServer(srv, identityHandler)
	identity.RegisterRoleServiceHTTPServer(srv, identityHandler)
	// token.RegisterTokenServiceHTTPServer(srv, tokenHandler)

	return srv
}
