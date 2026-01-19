package server

import (
	"crypto/tls"

	"github.com/dizzrt/dauth/api/gen/auth"
	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/api/gen/sp"
	"github.com/dizzrt/dauth/api/gen/token"
	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/dauth/internal/handler"
	"github.com/dizzrt/ellie/log"
	"github.com/dizzrt/ellie/middleware/tracing"
	"github.com/dizzrt/ellie/transport/grpc"
)

func NewGRPCServer(c *conf.AppConfig, logger log.LogWriter, identityHandler *handler.IdentityHandler, tokenHandler *handler.TokenHandler, spHandler *handler.ServiceProviderHandler, authHandler *handler.AuthHandler) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			tracing.UnaryServerInterceptor(),
		),
	}

	serverConf := c.Server.GRPC
	if serverConf.Addr != "" {
		opts = append(opts, grpc.Address(serverConf.Addr))
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

		opts = append(opts, grpc.TLSConfig(tlsConfig))
	} else {
		panic("dauth must run with TLS enabled. Please configure both server.grpc.tls.cert_path and server.grpc.tls.key_path in the configuration file")
	}

	srv := grpc.NewServer(opts...)
	identity.RegisterUserServiceServer(srv, identityHandler)
	identity.RegisterRoleServiceServer(srv, identityHandler)
	token.RegisterTokenServiceServer(srv, tokenHandler)
	sp.RegisterServiceProviderServiceServer(srv, spHandler)
	auth.RegisterAuthServiceServer(srv, authHandler)

	return srv
}
