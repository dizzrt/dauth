package handler

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/token"
	"github.com/dizzrt/dauth/internal/application"
	"github.com/dizzrt/ellie/errors"
)

var _ token.TokenServiceServer = (*TokenHandler)(nil)

type TokenHandler struct {
	token.UnimplementedTokenServiceServer

	tokenApp application.TokenApplication
}

func NewTokenHandler(tokenApp application.TokenApplication) *TokenHandler {
	return &TokenHandler{
		tokenApp: tokenApp,
	}
}

func (handler *TokenHandler) IssueSSOToken(ctx context.Context, req *token.IssueSSOTokenRequest) (*token.IssueSSOTokenResponse, error) {
	return errors.WrapGRPCResponse(handler.tokenApp.IssueSSOToken(ctx, req))
}

func (handler *TokenHandler) Issue(ctx context.Context, req *token.IssueRequest) (*token.IssueResponse, error) {
	return errors.WrapGRPCResponse(handler.tokenApp.Issue(ctx, req))
}

func (handler *TokenHandler) Validate(ctx context.Context, req *token.ValidateRequest) (*token.ValidateResponse, error) {
	return errors.WrapGRPCResponse(handler.tokenApp.Validate(ctx, req))
}

func (handler *TokenHandler) Revoke(ctx context.Context, req *token.RevokeRequest) (*token.RevokeResponse, error) {
	return errors.WrapGRPCResponse(handler.tokenApp.Revoke(ctx, req))
}
