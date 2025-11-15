package rpc

import (
	api_base "github.com/dizzrt/dauth/api/gen/base"
	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/infra/rpc/base"
)

const SuccessCode uint32 = 10000
const SuccessMessage string = "OK"

func NewBaseResp(code uint32, message string) *api_base.BaseResp {
	return &api_base.BaseResp{
		Code:    code,
		Message: message,
	}
}

func Success() *api_base.BaseResp {
	return NewBaseResp(SuccessCode, SuccessMessage)
}

func UserServiceClient() identity.UserServiceClient {
	return base.GetClient(base.CK_DAUTH_IDENTITY_USER).(identity.UserServiceClient)
}
