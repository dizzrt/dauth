package identity

import (
	"github.com/dizzrt/dauth/internal/domain/identity/biz"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	biz.NewUserBiz,
	biz.NewRoleBiz,
)
