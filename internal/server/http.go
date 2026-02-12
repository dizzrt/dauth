package server

import (
	"crypto/tls"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/dauth/internal/handler"
	"github.com/dizzrt/dauth/internal/server/middleware"
	"github.com/dizzrt/ellie/log"
	"github.com/dizzrt/ellie/middleware/tracing"
	"github.com/dizzrt/ellie/transport/http"
)

func NewHTTPServer(c *conf.AppConfig, logger log.LogWriter, identityHandler *handler.IdentityHandler) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			tracing.TracingMiddleware(),
			middleware.JwtAuthMiddleware(),
		),
	}

	serverConf := c.Server.HTTP
	if serverConf.Addr != "" {
		opts = append(opts, http.Address(serverConf.Addr))
	}

	if serverConf.TLSConfig.CertPath != "" && serverConf.TLSConfig.KeyPath != "" {
		tlsConfig := &tls.Config{
			MinVersion: tls.VersionTLS12,
		}

		// load server certificate and key
		cert, err := tls.LoadX509KeyPair(serverConf.TLSConfig.CertPath, serverConf.TLSConfig.KeyPath)
		if err != nil {
			panic(err)
		}

		tlsConfig.Certificates = []tls.Certificate{cert}
		if c.ENV == "dev" {
			// skip tls verify in dev env
			tlsConfig.InsecureSkipVerify = true
		}

		opts = append(opts, http.TLSConfig(tlsConfig))
	} else {
		panic("dauth must run with TLS enabled. Please configure both server.http.tls.cert_path and server.http.tls.key_path in the configuration file")
	}

	srv := http.NewServer(opts...)
	identity.RegisterUserServiceHTTPServer(srv, identityHandler)

	return srv
}
