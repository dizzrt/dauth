package dauth

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/infra/rpc"
	"github.com/dizzrt/ellie/errors"
)

func GetUser(ctx context.Context, uid uint32) (*identity.GetUserResponse, error) {
	req := &identity.GetUserRequest{
		Id: uid,
	}

	return A(rpc.UserServiceClient().GetUser(ctx, req))
}

func A[T any](v T, err error) (T, error) {
	return v, errors.Unmarshal(err)
}
