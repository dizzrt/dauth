package application

import (
	"github.com/dizzrt/dauth/internal/domain/identity/biz"
)

var _ IdentityApplication = (*identityApplication)(nil)

type IdentityApplication interface {
	// Authenticate(context.Context, *identity.AuthenticateRequest) (*identity.AuthenticateResponse, error)
}

type identityApplication struct {
	userBiz biz.UserBiz
	roleBiz biz.RoleBiz
}

func NewIdentityApplication(userBiz biz.UserBiz, roleBiz biz.RoleBiz) IdentityApplication {
	return &identityApplication{
		userBiz: userBiz,
		roleBiz: roleBiz,
	}
}
