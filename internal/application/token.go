package application

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/token"
	"github.com/dizzrt/dauth/internal/domain/token/biz"
)

var _ TokenApplication = (*tokenApplication)(nil)

type TokenApplication interface {
	Issue(context.Context, *token.IssueRequest) (*token.IssueResponse, error)
	Validate(context.Context, *token.ValidateRequest) (*token.ValidateResponse, error)
	Revoke(context.Context, *token.RevokeRequest) (*token.RevokeResponse, error)
}

type tokenApplication struct {
	tokenBiz biz.TokenBiz
}

func NewTokenApplication(tokenBiz biz.TokenBiz) TokenApplication {
	return &tokenApplication{
		tokenBiz: tokenBiz,
	}
}

func (app *tokenApplication) Issue(ctx context.Context, req *token.IssueRequest) (*token.IssueResponse, error) {
	return nil, nil
}

func (app *tokenApplication) Validate(ctx context.Context, req *token.ValidateRequest) (*token.ValidateResponse, error) {
	return nil, nil
}

func (app *tokenApplication) Revoke(ctx context.Context, req *token.RevokeRequest) (*token.RevokeResponse, error) {
	return nil, nil
}
