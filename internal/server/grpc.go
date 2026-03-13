package server

import (
	"crypto/tls"

	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/dauth/internal/handler"
	"github.com/dizzrt/ellie/log"
	"github.com/dizzrt/ellie/middleware/tracing"
	"github.com/dizzrt/ellie/transport/grpc"
)

func NewGRPCServer(c *conf.AppConfig, logger log.LogWriter, registrar *handler.ServiceRegistrar) *grpc.Server {
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
	registrar.Register(srv)

	return srv
}
