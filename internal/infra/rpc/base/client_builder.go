package base

import (
	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/api/gen/sp"
	"google.golang.org/grpc"
)

func NewUserServiceClient(conn *grpc.ClientConn) (any, error) {
	return identity.NewUserServiceClient(conn), nil
}

func NewServiceProviderServiceClient(conn *grpc.ClientConn) (any, error) {
	return sp.NewServiceProviderServiceClient(conn), nil
}
