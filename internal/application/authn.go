package application

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/authn"
)

var _ AuthnApplication = (*authnApplication)(nil)

type AuthnApplication interface {
	Login(ctx context.Context, req *authn.LoginRequest) (*authn.LoginResponse, error)
	Logout(ctx context.Context, req *authn.LogoutRequest) (*authn.LogoutResponse, error)
}

type authnApplication struct {
}

func NewAuthnApplication() AuthnApplication {
	return &authnApplication{}
}

func (app *authnApplication) Login(ctx context.Context, req *authn.LoginRequest) (*authn.LoginResponse, error) {
	return nil, nil
}

func (app *authnApplication) Logout(ctx context.Context, req *authn.LogoutRequest) (*authn.LogoutResponse, error) {
	return nil, nil
}
