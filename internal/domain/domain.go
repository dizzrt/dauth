package domain

import (
	identity_biz "github.com/dizzrt/dauth/internal/domain/identity/biz"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	identity_biz.NewUserBiz,
	identity_biz.NewRoleBiz,
)
