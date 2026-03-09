package dauth

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/infra/rpc"
	"github.com/dizzrt/ellie/errors"
)

func GetUser(ctx context.Context, uid uint32) (*identity.GetUserResponse, error) {
	req := &identity.GetUserRequest{
		Uid: uid,
	}

	return errors.UnwrapGRPCResponse(rpc.UserServiceClient().GetUser(ctx, req))
}

func GetUserByName(ctx context.Context, username string) (*identity.GetUserByNameResponse, error) {
	req := &identity.GetUserByNameRequest{
		Username: username,
	}

	return errors.UnwrapGRPCResponse(rpc.UserServiceClient().GetUserByName(ctx, req))
}

func VerifyPassword(ctx context.Context, req *identity.VerifyPasswordRequest) (*identity.VerifyPasswordResponse, error) {
	return errors.UnwrapGRPCResponse(rpc.UserServiceClient().VerifyPassword(ctx, req))
}
