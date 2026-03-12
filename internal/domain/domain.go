package domain

import (
	authn_biz "github.com/dizzrt/dauth/internal/domain/authn/biz"
	identity_biz "github.com/dizzrt/dauth/internal/domain/identity/biz"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	identity_biz.NewUserBiz,
	authn_biz.NewAuthnBiz,
)
