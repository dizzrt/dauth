package application

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/domain/identity/biz"
)

var _ IdentityApplication = (*identityApplication)(nil)

type IdentityApplication interface {
	CreateUser(context.Context, *identity.CreateUserRequest) (*identity.CreateUserResponse, error)
	GetUser(context.Context, *identity.GetUserRequest) (*identity.GetUserResponse, error)
}

type identityApplication struct {
	userBiz biz.UserBiz
}

func NewIdentityApplication(userBiz biz.UserBiz) IdentityApplication {
	return &identityApplication{
		userBiz: userBiz,
	}
}

func (app *identityApplication) CreateUser(ctx context.Context, req *identity.CreateUserRequest) (*identity.CreateUserResponse, error) {
	return nil, nil
}

func (app *identityApplication) GetUser(ctx context.Context, req *identity.GetUserRequest) (*identity.GetUserResponse, error) {
	return nil, nil
}
