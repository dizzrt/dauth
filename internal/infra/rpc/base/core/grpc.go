package core

import (
	"sync"
	"time"

	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/ellie/contrib/registry/consul"
	"github.com/dizzrt/ellie/log"
	"github.com/dizzrt/ellie/middleware/tracing"
	"github.com/dizzrt/ellie/registry"
	transport_grpc "github.com/dizzrt/ellie/transport/grpc"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

type ClientBuilder func(*grpc.ClientConn) (any, error)

type Client struct {
	Key      string
	Endpoint string
	Builder  ClientBuilder
}

var (
	manager    *clientManager
	discoverer registry.Discovery
)

func init() {
	// init consul discovery
	ac := conf.GetAppConfig()
	client, err := api.NewClient(&api.Config{
		Address: ac.Registry.Addr,
	})

	if err != nil {
		panic(err)
	}

	discoverer = consul.New(client)

	// init conn manager
	manager = &clientManager{}
}

type clientManager struct {
	conns   sync.Map
	clients sync.Map
}

func (cm *clientManager) newConn(endpoint string) (*grpc.ClientConn, error) {
	conn, err := transport_grpc.DialInsecure(
		transport_grpc.WithEndpoint(endpoint),
		transport_grpc.WithDiscovery(discoverer),
		transport_grpc.WithPrintDiscoveryDebugLog(true),
		transport_grpc.WithTimeout(5*time.Second),
		transport_grpc.WithUnaryClientInterceptor(
			tracing.UnaryClientInterceptor(),
		),
	)

	if err != nil {
		log.Errorf("failed to create grpc conn for endpoint %s, err: %v", endpoint, err)
		return nil, err
	}

	cm.conns.Store(endpoint, conn)
	return conn, err
}

func (cm *clientManager) getConn(endpoint string) (*grpc.ClientConn, error) {
	temp, ok := cm.conns.Load(endpoint)
	if !ok {
		return cm.newConn(endpoint)
	}

	conn, ok := temp.(*grpc.ClientConn)
	if !ok {
		return cm.newConn(endpoint)
	}

	return conn, nil
}

func (cm *clientManager) getClient(key string) any {
	client, ok := cm.clients.Load(key)
	if !ok {
		return nil
	}

	return client
}

func (cm *clientManager) registerClient(key string, client any) {
	cm.clients.Store(key, client)
}

func GetClient(key string) any {
	return manager.getClient(key)
}

func NewClients(clients ...Client) {
	existKeys := make(map[string]struct{}, len(clients))
	for _, cli := range clients {
		if _, ok := existKeys[cli.Key]; ok {
			panic("duplicate client key")
		}

		conn, err := manager.getConn(cli.Endpoint)
		if err != nil {
			panic(err)
		}

		c, err := cli.Builder(conn)
		if err != nil {
			panic(err)
		}

		manager.registerClient(cli.Key, c)
		existKeys[cli.Key] = struct{}{}
	}
}
