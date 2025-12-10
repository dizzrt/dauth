package server

import (
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

	grpcServerConf := c.Server.GRPC
	if grpcServerConf.Addr != "" {
		opts = append(opts, grpc.Address(grpcServerConf.Addr))
	}

	srv := grpc.NewServer(opts...)
	identity.RegisterUserServiceServer(srv, identityHandler)
	identity.RegisterRoleServiceServer(srv, identityHandler)
	token.RegisterTokenServiceServer(srv, tokenHandler)
	sp.RegisterServiceProviderServiceServer(srv, spHandler)
	auth.RegisterAuthServiceServer(srv, authHandler)

	return srv
}
