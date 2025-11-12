package dauth

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/infra/rpc"
)

var (
	// client *UserClient
	client identity.UserServiceClient
)

// type UserClient struct {
// 	rpc.GRPCBaseClient
// 	stub identity.UserServiceClient
// }

func init() {
	conn, err := rpc.NewGRPCBaseClient("discovery:///dauth")
	if err != nil {
		panic(err)
	}

	client = identity.NewUserServiceClient(conn)
}

func GetUser(ctx context.Context, req *identity.GetUserRequest) (*identity.GetUserResponse, error) {
	return client.GetUser(ctx, req)
}
