package domain

import (
	client_biz "github.com/dizzrt/dauth/internal/domain/client/biz"
	identity_biz "github.com/dizzrt/dauth/internal/domain/identity/biz"
	token_biz "github.com/dizzrt/dauth/internal/domain/token/biz"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	identity_biz.NewUserBiz,
	identity_biz.NewRoleBiz,
	token_biz.NewTokenBiz,
	client_biz.NewClientBiz,
)
