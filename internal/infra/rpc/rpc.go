package rpc

import (
	"github.com/dizzrt/dauth/api/gen/authn"
	api_common "github.com/dizzrt/dauth/api/gen/common"
	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/infra/rpc/base"
)

const SuccessCode uint32 = 10000
const SuccessMessage string = "OK"

func NewBaseResp(code uint32, message string) *api_common.BaseResp {
	return &api_common.BaseResp{
		Code:    code,
		Message: message,
	}
}

func Success() *api_common.BaseResp {
	return NewBaseResp(SuccessCode, SuccessMessage)
}

func UserServiceClient() identity.UserServiceClient {
	return base.GetClient(base.CK_DAUTH_IDENTITY_USER).(identity.UserServiceClient)
}

func AuthnServiceClient() authn.AuthnServiceClient {
	return base.GetClient(base.CK_DAUTH_AUTHN).(authn.AuthnServiceClient)
}
