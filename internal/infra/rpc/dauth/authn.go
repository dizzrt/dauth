package dauth

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/authn"
	"github.com/dizzrt/dauth/internal/infra/rpc"
	"github.com/dizzrt/ellie/errors"
)

func CheckAuthnStatus(ctx context.Context, token string) (*authn.CheckAuthnStatusResponse, error) {
	req := &authn.CheckAuthnStatusRequest{
		Token: token,
	}

	return errors.UnwrapGRPCResponse(rpc.AuthnServiceClient().CheckAuthnStatus(ctx, req))
}
