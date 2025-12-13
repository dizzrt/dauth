package middleware

import (
	"slices"

	"github.com/dizzrt/dauth/api/gen/token"
	"github.com/dizzrt/dauth/internal/infra/rpc/dauth"
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

		if err != nil || resp.GetToken().Uid == 0 {
			unauthorized(ctx)
			return
		}

		ctx.Set("uid", resp.GetToken().Uid)
		ctx.Next()
	}
}
