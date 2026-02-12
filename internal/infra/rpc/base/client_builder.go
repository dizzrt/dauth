package base

import (
	"github.com/dizzrt/dauth/api/gen/identity"
	"google.golang.org/grpc"
)

func NewUserServiceClient(conn *grpc.ClientConn) (any, error) {
	return identity.NewUserServiceClient(conn), nil
}
