package ctxutil

import (
	"context"

	"github.com/dizzrt/ellie/log"
	"github.com/gin-gonic/gin"
)

const (
	_CTXUTIL_CONST_KEY_UID = "_ctxutil_const_key_uid"
)

func SetUid[T any](ctx T, uid uint32) T {
	switch v := any(ctx).(type) {
	case *gin.Context:
		ctx = any(setUidForGinCtx(v, uid)).(T)
	case context.Context:
		ctx = any(setUidForCtx(v, uid)).(T)
	}

	return ctx
}

func setUidForGinCtx(ctx *gin.Context, uid uint32) *gin.Context {
	ctx.Set(_CTXUTIL_CONST_KEY_UID, uid)

	rctx := ctx.Request.Context()
	rctx = setUidForCtx(rctx, uid)
	ctx.Request = ctx.Request.WithContext(rctx)

	return ctx
}

func setUidForCtx(ctx context.Context, uid uint32) context.Context {
	ctx = context.WithValue(ctx, _CTXUTIL_CONST_KEY_UID, uid)
	return ctx
}

func GetUid[T any](ctx T) uint32 {
	var uid uint32 = 0
	switch v := any(ctx).(type) {
	case *gin.Context:
		uid = getUidFromGinCtx(v)
	case context.Context:
		uid = getUidFromCtx(v)
	}

	return uid
}

func getUidFromGinCtx(ctx *gin.Context) uint32 {
	return ctx.GetUint32(_CTXUTIL_CONST_KEY_UID)
}

func getUidFromCtx(ctx context.Context) uint32 {
	v := ctx.Value(_CTXUTIL_CONST_KEY_UID)
	if v == nil {
		return 0
	}

	uid, ok := v.(uint32)
	if !ok {
		log.CtxWarnf(ctx, "assert uid failed, uid: %v", uid)
		return 0
	}

	return uid
}
