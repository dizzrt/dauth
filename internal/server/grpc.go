package server

import (
	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/dauth/internal/handler"
	"github.com/dizzrt/ellie/log"
	"github.com/dizzrt/ellie/middleware/tracing"
	"github.com/dizzrt/ellie/transport/grpc"
)

func NewGRPCServer(c *conf.Bootstrap, logger log.LogWriter, identityHandler *handler.IdentityHandler) *grpc.Server {
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

	return srv
}
