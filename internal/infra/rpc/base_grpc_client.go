package rpc

import (
	"time"

	"github.com/dizzrt/ellie/contrib/registry/consul"
	"github.com/dizzrt/ellie/registry"
	transport_grpc "github.com/dizzrt/ellie/transport/grpc"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

var (
	discoverer registry.Discovery
)

func init() {
	client, err := api.NewClient(&api.Config{
		Address: "infra.dauth.com:8500",
	})

	if err != nil {
		panic(err)
	}

	discoverer = consul.New(client)
}

type GRPCBaseClient struct {
	conn *grpc.ClientConn
}

func NewGRPCBaseClient(endpoint string) (*grpc.ClientConn, error) {
	conn, err := transport_grpc.DialInsecure(
		transport_grpc.WithEndpoint(endpoint),
		transport_grpc.WithDiscovery(discoverer),
		transport_grpc.WithPrintDiscoveryDebugLog(true),
		transport_grpc.WithTimeout(5*time.Second),
	)

	return conn, err
}
