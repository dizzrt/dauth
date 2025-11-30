package dauth

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/client"
	"github.com/dizzrt/dauth/internal/infra/rpc"
)

func GetClient(ctx context.Context, clientID uint32) (*client.GetClientResponse, error) {
	req := &client.GetClientRequest{
		ClientId: clientID,
	}

	return rpc.ClientServiceClient().GetClient(ctx, req)
}

func ValidateClient(ctx context.Context, clientID uint32, scope string) (*client.ValidateClientResponse, error) {
	req := &client.ValidateClientRequest{
		ClientId: clientID,
		Scope:    scope,
	}

	return rpc.ClientServiceClient().ValidateClient(ctx, req)
}
