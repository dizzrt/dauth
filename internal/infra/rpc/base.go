package rpc

import "github.com/dizzrt/dauth/api/gen/base"

const SuccessCode uint32 = 10000
const SuccessMessage string = "OK"

func NewBaseResp(code uint32, message string) *base.BaseResp {
	return &base.BaseResp{
		Code:    code,
		Message: message,
	}
}

func Success() *base.BaseResp {
	return NewBaseResp(SuccessCode, SuccessMessage)
}
