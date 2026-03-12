package application

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/authn"
	"github.com/dizzrt/dauth/api/gen/errdef"
	"github.com/dizzrt/dauth/internal/domain/authn/biz"
	"github.com/dizzrt/dauth/internal/domain/authn/dto"
	"github.com/dizzrt/ellie/log"
)

var _ AuthnApplication = (*authnApplication)(nil)

type AuthnApplication interface {
	Login(ctx context.Context, req *authn.LoginRequest) (*authn.LoginResponse, error)
	Logout(ctx context.Context, req *authn.LogoutRequest) (*authn.LogoutResponse, error)
}

type authnApplication struct {
	authnBiz biz.AuthnBiz
}

func NewAuthnApplication(authnBiz biz.AuthnBiz) AuthnApplication {
	return &authnApplication{authnBiz: authnBiz}
}

func (app *authnApplication) Login(ctx context.Context, req *authn.LoginRequest) (*authn.LoginResponse, error) {
	account := req.GetAccount()
	if account == "" {
		return nil, errdef.InvalidArgument().WithMessage("account can not be empty")
	}

	password := req.GetPassword()
	if password == "" {
		return nil, errdef.InvalidArgument().WithMessage("password can not be empty")
	}

	bizResp, err := app.authnBiz.Login(ctx, account, password)
	if err != nil {
		log.CtxErrorf(ctx, "login failed, account: %s, err: %v", account, err)
		return nil, err
	}

	resp := &authn.LoginResponse{}
	if bizResp.Status == dto.LoginStatusSuccess {
		resp.Status = authn.AuthnAttemptStatus_SUCCESS
		resp.Token = &bizResp.Token
	} else {
		resp.Status = authn.AuthnAttemptStatus_FAILED
	}

	return resp, nil
}

func (app *authnApplication) Logout(ctx context.Context, req *authn.LogoutRequest) (*authn.LogoutResponse, error) {
	return nil, nil
}
