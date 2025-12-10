package dauth

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/sp"
	"github.com/dizzrt/dauth/internal/infra/rpc"
)

func GetServiceProvider(ctx context.Context, spID uint32) (*sp.GetServiceProviderResponse, error) {
	req := &sp.GetServiceProviderRequest{
		SpId: spID,
	}

	return rpc.ServiceProviderServiceClient().GetServiceProvider(ctx, req)
}

func ValidateServiceProvider(ctx context.Context, spID uint32, scope string) (*sp.ValidateServiceProviderResponse, error) {
	req := &sp.ValidateServiceProviderRequest{
		SpId:  spID,
		Scope: scope,
	}

	return rpc.ServiceProviderServiceClient().ValidateServiceProvider(ctx, req)
}
