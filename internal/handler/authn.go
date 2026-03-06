package handler

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/authn"
	"github.com/dizzrt/dauth/internal/application"
	"github.com/dizzrt/ellie/errors"
)

var _ authn.AuthnServiceServer = (*AuthnHandler)(nil)

type AuthnHandler struct {
	authn.UnimplementedAuthnServiceServer

	authnApp application.AuthnApplication
}

func NewAuthnHandler(authnApp application.AuthnApplication) *AuthnHandler {
	return &AuthnHandler{authnApp: authnApp}
}

func (handler *AuthnHandler) Login(ctx context.Context, req *authn.LoginRequest) (*authn.LoginResponse, error) {
	return errors.WrapGRPCResponse(handler.authnApp.Login(ctx, req))
}

func (handler *AuthnHandler) Logout(ctx context.Context, req *authn.LogoutRequest) (*authn.LogoutResponse, error) {
	return errors.WrapGRPCResponse(handler.authnApp.Logout(ctx, req))
}
