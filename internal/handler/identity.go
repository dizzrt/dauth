package handler

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/application"
	"github.com/dizzrt/ellie/errors"
)

var _ identity.UserServiceServer = (*IdentityHandler)(nil)

type IdentityHandler struct {
	identity.UnimplementedUserServiceServer

	identityApp application.IdentityApplication
}

func NewIdentityHandler(identityApp application.IdentityApplication) *IdentityHandler {
	return &IdentityHandler{
		identityApp: identityApp,
	}
}

func (handler *IdentityHandler) CreateUser(ctx context.Context, req *identity.CreateUserRequest) (*identity.CreateUserResponse, error) {
	return errors.WrapGRPCResponse(handler.identityApp.CreateUser(ctx, req))
}

func (handler *IdentityHandler) GetUser(ctx context.Context, req *identity.GetUserRequest) (*identity.GetUserResponse, error) {
	return errors.WrapGRPCResponse(handler.identityApp.GetUser(ctx, req))
}

func (handler *IdentityHandler) ListUsers(ctx context.Context, req *identity.ListUsersRequest) (*identity.ListUsersResponse, error) {
	return errors.WrapGRPCResponse(handler.identityApp.ListUsers(ctx, req))
}

func (handler *IdentityHandler) UpdateUserStatus(ctx context.Context, req *identity.UpdateUserStatusRequest) (*identity.UpdateUserStatusResponse, error) {
	return errors.WrapGRPCResponse(handler.identityApp.UpdateUserStatus(ctx, req))
}
