package foundation

import (
	"time"

	"github.com/dizzrt/dauth/internal/conf"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	IsCluster     bool
	client        *redis.Client
	clusterClient *redis.ClusterClient
}

func NewRedisClient(ac *conf.AppConfig) (*RedisClient, func(), error) {
	client := redis.NewClient(&redis.Options{
		Addr:         ac.Redis.Addr,
		Password:     ac.Redis.Password,
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 5,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	})

	cli := &RedisClient{
		IsCluster:     false,
		client:        client,
		clusterClient: nil,
	}

	return cli, func() {
		_ = client.Close()
	}, nil
}

func (rc *RedisClient) Cmdable() redis.Cmdable {
	if rc.IsCluster {
		return rc.clusterClient
	}

	return rc.client
}

func (rc *RedisClient) Close() error {
	if rc.IsCluster {
		return rc.clusterClient.Close()
	}

	return rc.client.Close()
}
