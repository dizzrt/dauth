package middleware

import (
	"slices"

	"github.com/dizzrt/dauth/api/gen/errdef"
	"github.com/dizzrt/dauth/api/gen/token"
	"github.com/dizzrt/dauth/internal/infra/rpc/dauth"
	"github.com/dizzrt/dauth/internal/infra/utils/ctxutil"
	"github.com/dizzrt/ellie/errors"
	"github.com/dizzrt/ellie/log"
	"github.com/dizzrt/ellie/transport/http"
	"github.com/gin-gonic/gin"
)

var authWhiteList = []string{
	"/identity/user/login",
}

func unauthorized(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"data": nil, "message": "unauthorized", "status": http.StatusUnauthorized})
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if slices.Contains(authWhiteList, ctx.FullPath()) {
			ctx.Next()
			return
		}

		tokenStr := ctx.Request.Header.Get("Authorization")
		if tokenStr == "" {
			unauthorized(ctx)
			return
		}

		resp, err := dauth.ValidateToken(ctx.Request.Context(), &token.ValidateRequest{
			Token: tokenStr,
			Type:  token.Token_TokenType_SSO,
		})

		if err != nil {
			if !errors.Is(err, errdef.TokenExpired()) && !errors.Is(err, errdef.TokenRevoked()) && !errors.Is(err, errdef.TokenInvalid()) {
				log.CtxErrorf(ctx, "validate token failed, token: %s, err: %v", tokenStr, err)
			}

			unauthorized(ctx)
			return
		}

		uid := resp.GetToken().GetUid()
		if uid == 0 {
			unauthorized(ctx)
			return
		}

		ctxutil.SetUid(ctx, uid) // inject uid to context
		ctx.Next()
	}
}
