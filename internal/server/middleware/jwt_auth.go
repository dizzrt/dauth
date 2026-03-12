package middleware

import (
	"slices"

	"github.com/dizzrt/dauth/api/gen/authn"
	"github.com/dizzrt/dauth/internal/infra/rpc/dauth"
	"github.com/dizzrt/dauth/internal/infra/utils/ctxutil"
	"github.com/dizzrt/ellie/log"
	"github.com/dizzrt/ellie/transport/http"
	"github.com/gin-gonic/gin"
)

func unauthorized(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"data": nil, "message": "unauthorized", "status": http.StatusUnauthorized})
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		if wl, ok := authWhiteList[method]; ok {
			if slices.Contains(wl, ctx.FullPath()) {
				ctx.Next()
				return
			}
		}

		tokenStr := ctx.Request.Header.Get("Authorization")
		if tokenStr == "" {
			unauthorized(ctx)
			return
		}

		resp, err := dauth.CheckAuthnStatus(ctx.Request.Context(), tokenStr)
		if err != nil {
			log.CtxErrorf(ctx, "check authn status failed, token: %s, err: %v", tokenStr, err)

			unauthorized(ctx)
			return
		}

		if resp.GetStatus() != authn.AuthnStatus_VALID {
			unauthorized(ctx)
			return
		}

		uid := resp.GetUid()
		sid := resp.GetSid()
		if uid == 0 || sid == "" {
			unauthorized(ctx)
			return
		}

		// inject some context values
		ctxutil.SetUid(ctx, uid)
		ctxutil.SetSessionID(ctx, sid)
		ctx.Next()
	}
}
