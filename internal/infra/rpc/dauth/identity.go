package dauth

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/infra/rpc"
)

func GetUser(ctx context.Context, uid uint32) (*identity.GetUserResponse, error) {
	req := &identity.GetUserRequest{
		Id: uid,
	}

	return rpc.UserServiceClient().GetUser(ctx, req)
}
