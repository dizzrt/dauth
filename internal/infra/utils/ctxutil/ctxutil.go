package ctxutil

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ctxKey string

const (
	_CTXUTIL_CONST_KEY_UID ctxKey = "_ctxutil_const_key_uid"
	_CTXUTIL_CONST_KEY_SID ctxKey = "_ctxutil_const_key_sid"
)

func SetUid[T any](ctx T, uid uint32) T {
	return setWithKey(ctx, _CTXUTIL_CONST_KEY_UID, uid)
}

func GetUid[T any](ctx T) uint32 {
	return getWithKey[T, uint32](ctx, _CTXUTIL_CONST_KEY_UID)
}

func SetSessionID[T any](ctx T, sid string) T {
	return setWithKey(ctx, _CTXUTIL_CONST_KEY_SID, sid)
}

func GetSessionID[T any](ctx T) string {
	return getWithKey[T, string](ctx, _CTXUTIL_CONST_KEY_SID)
}

func getWithKey[T, S any](ctx T, key ctxKey) S {
	var temp any
	switch v := any(ctx).(type) {
	case *gin.Context:
		temp = getWithKeyFromGinCtx(v, key)
	case context.Context:
		temp = getWithKeyFromCtx(v, key)
	}

	var s S
	if temp == nil {
		return s
	}

	if v, ok := temp.(S); ok {
		return v
	}

	return s
}

func getWithKeyFromGinCtx(ctx *gin.Context, key ctxKey) any {
	v, _ := ctx.Get(key)
	return v
}

func getWithKeyFromCtx(ctx context.Context, key ctxKey) any {
	return ctx.Value(key)
}

func setWithKey[T any](ctx T, key ctxKey, value any) T {
	switch v := any(ctx).(type) {
	case *gin.Context:
		ctx = any(setWithKeyForGinCtx(v, key, value)).(T)
	case context.Context:
		ctx = any(setWithKeyForCtx(v, key, value)).(T)
	}

	return ctx
}

func setWithKeyForGinCtx(ctx *gin.Context, key ctxKey, value any) *gin.Context {
	ctx.Set(key, value)

	rctx := ctx.Request.Context()
	rctx = setWithKeyForCtx(rctx, key, value)
	ctx.Request = ctx.Request.WithContext(rctx)

	return ctx
}

func setWithKeyForCtx(ctx context.Context, key ctxKey, value any) context.Context {
	ctx = context.WithValue(ctx, key, value)
	return ctx
}
