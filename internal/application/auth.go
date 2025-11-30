package application

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/auth"
	"github.com/dizzrt/dauth/api/gen/client"
	"github.com/dizzrt/dauth/api/gen/errdef"
	"github.com/dizzrt/dauth/internal/domain/auth/biz"
	"github.com/dizzrt/dauth/internal/infra/rpc"
	"github.com/dizzrt/dauth/internal/infra/rpc/dauth"
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
	scope := req.GetScope()
	UserID := req.GetUserId()
	clientID := req.GetClientId()
	redirectURI := req.GetRedirectUri()
	if UserID == 0 || clientID == 0 || redirectURI == "" {
		return nil, errdef.InvalidParams()
	}

	// verify client
	resp, err := dauth.GetClient(ctx, clientID)
	if err != nil {
		return nil, err
	}

	cli := resp.GetClient()
	if cli == nil || cli.Status != client.Client_ACTIVE {
		return nil, errdef.AuthInvalidClient()
	}

	if cli.GetRedirectUri() != redirectURI {
		return nil, errdef.AuthInvalidRedirectURI()
	}

	// TODO check scope

	code, err := app.authBiz.GenerateAuthorizationCode(ctx, UserID, clientID, redirectURI, scope)
	if err != nil {
		return nil, err
	}

	return &auth.GenerateAuthorizationCodeResponse{
		Code:     code,
		BaseResp: rpc.Success(),
	}, nil
}

func (app *authApplication) ExchangeToken(ctx context.Context, req *auth.ExchangeTokenRequest) (*auth.ExchangeTokenResponse, error) {
	return nil, nil
}
