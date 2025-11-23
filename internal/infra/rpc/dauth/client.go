package dauth

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/client"
	"github.com/dizzrt/dauth/internal/infra/rpc"
)

func ValidateClient(ctx context.Context, clientID uint32, scope string) (*client.ValidateResponse, error) {
	req := &client.ValidateRequest{
		ClientId: clientID,
		Scope:    scope,
	}

	return rpc.ClientServiceClient().Validate(ctx, req)
}
