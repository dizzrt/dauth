package dauth

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/token"
	"github.com/dizzrt/dauth/internal/infra/rpc"
)

func IssueSSOToken(ctx context.Context, uid uint32) (*token.IssueSSOTokenResponse, error) {
	req := &token.IssueSSOTokenRequest{
		Uid: uid,
	}

	return rpc.TokenServiceClient().IssueSSOToken(ctx, req)
}

func ValidateToken(ctx context.Context, req *token.ValidateRequest) (*token.ValidateResponse, error) {
	return rpc.TokenServiceClient().Validate(ctx, req)
}
