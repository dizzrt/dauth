package handler

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/auth"
	"github.com/dizzrt/dauth/internal/application"
)

var _ auth.AuthServiceServer = (*AuthHandler)(nil)

type AuthHandler struct {
	auth.UnimplementedAuthServiceServer

	authApp application.AuthApplication
}

func NewAuthHandler(authApp application.AuthApplication) *AuthHandler {
	return &AuthHandler{
		authApp: authApp,
	}
}

func (handler *AuthHandler) GenerateAuthorizationCode(ctx context.Context, req *auth.GenerateAuthorizationCodeRequest) (*auth.GenerateAuthorizationCodeResponse, error) {
	return handler.authApp.GenerateAuthorizationCode(ctx, req)
}

func (handler *AuthHandler) ExchangeToken(ctx context.Context, req *auth.ExchangeTokenRequest) (*auth.ExchangeTokenResponse, error) {
	return handler.authApp.ExchangeToken(ctx, req)
}
