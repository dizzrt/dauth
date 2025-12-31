package dauth

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/sp"
	"github.com/dizzrt/dauth/internal/infra/rpc"
	"github.com/dizzrt/ellie/errors"
)

func GetServiceProvider(ctx context.Context, spID uint32) (*sp.GetServiceProviderResponse, error) {
	req := &sp.GetServiceProviderRequest{
		SpId: spID,
	}

	return errors.UnwrapGRPCResponse(rpc.ServiceProviderServiceClient().GetServiceProvider(ctx, req))
}

func ValidateServiceProvider(ctx context.Context, spID uint32, scope string) (*sp.ValidateServiceProviderResponse, error) {
	req := &sp.ValidateServiceProviderRequest{
		SpId:  spID,
		Scope: scope,
	}

	return errors.UnwrapGRPCResponse(rpc.ServiceProviderServiceClient().ValidateServiceProvider(ctx, req))
}
