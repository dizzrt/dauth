package domain

import (
	auth_biz "github.com/dizzrt/dauth/internal/domain/auth/biz"
	identity_biz "github.com/dizzrt/dauth/internal/domain/identity/biz"
	sp_biz "github.com/dizzrt/dauth/internal/domain/sp/biz"
	token_biz "github.com/dizzrt/dauth/internal/domain/token/biz"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	identity_biz.NewUserBiz,
	identity_biz.NewRoleBiz,
	token_biz.NewTokenBiz,
	sp_biz.NewServiceProviderBiz,
	auth_biz.NewAuthBiz,
)
