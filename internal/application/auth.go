package application

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/auth"
	"github.com/dizzrt/dauth/internal/domain/auth/biz"
)

var _ AuthApplication = (*authApplication)(nil)

type AuthApplication interface {
	GenerateAuthorizationCode(ctx context.Context, req *auth.GenerateAuthorizationCodeRequest) (*auth.GenerateAuthorizationCodeResponse, error)
	ExchangeToken(ctx context.Context, req *auth.ExchangeTokenRequest) (*auth.ExchangeTokenResponse, error)
}

type authApplication struct {
	authBiz biz.AuthBiz
}

func NewAuthApplication(authBiz biz.AuthBiz) AuthApplication {
	return &authApplication{
		authBiz: authBiz,
	}
}

func (app *authApplication) GenerateAuthorizationCode(ctx context.Context, req *auth.GenerateAuthorizationCodeRequest) (*auth.GenerateAuthorizationCodeResponse, error) {
	return nil, nil
}

func (app *authApplication) ExchangeToken(ctx context.Context, req *auth.ExchangeTokenRequest) (*auth.ExchangeTokenResponse, error) {
	return nil, nil
}
