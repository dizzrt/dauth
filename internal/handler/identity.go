package handler

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/application"
	"github.com/dizzrt/ellie/errors"
)

var _ identity.UserServiceServer = (*IdentityHandler)(nil)
var _ identity.RoleServiceServer = (*IdentityHandler)(nil)

type IdentityHandler struct {
	identity.UnimplementedUserServiceServer
	identity.UnimplementedRoleServiceServer

	identityApp application.IdentityApplication
}

func NewIdentityHandler(identityApp application.IdentityApplication) *IdentityHandler {
	return &IdentityHandler{
		identityApp: identityApp,
	}
}

// region user service

func (handler *IdentityHandler) Login(ctx context.Context, req *identity.LoginRequest) (*identity.LoginResponse, error) {
	return errors.WrapGRPCResponse(handler.identityApp.Login(ctx, req))
}

func (handler *IdentityHandler) Authenticate(ctx context.Context, req *identity.AuthenticateRequest) (*identity.AuthenticateResponse, error) {
	return errors.WrapGRPCResponse(handler.identityApp.Authenticate(ctx, req))
}

func (handler *IdentityHandler) CreateUser(ctx context.Context, req *identity.CreateUserRequest) (*identity.CreateUserResponse, error) {
	return errors.WrapGRPCResponse(handler.identityApp.CreateUser(ctx, req))
}

func (handler *IdentityHandler) GetUser(ctx context.Context, req *identity.GetUserRequest) (*identity.GetUserResponse, error) {
	return errors.WrapGRPCResponse(handler.identityApp.GetUser(ctx, req))
}

func (handler *IdentityHandler) UpdateUserStatus(ctx context.Context, req *identity.UpdateUserStatusRequest) (*identity.UpdateUserStatusResponse, error) {
	return errors.WrapGRPCResponse(handler.identityApp.UpdateUserStatus(ctx, req))
}

func (handler *IdentityHandler) UpdateUserPassword(ctx context.Context, req *identity.UpdateUserPasswordRequest) (*identity.UpdateUserPasswordResponse, error) {
	return errors.WrapGRPCResponse(handler.identityApp.UpdateUserPassword(ctx, req))
}

// endregion user service

// region role service

// TODO: implement role service

// endregion role service
