package handler

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/application"
)

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

func (handler *IdentityHandler) Authenticate(ctx context.Context, req *identity.AuthenticateRequest) (*identity.AuthenticateResponse, error) {
	return nil, nil
	// return handler.identityApp.Authenticate(ctx, req)
}
