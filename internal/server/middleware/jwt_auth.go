package middleware

import (
	"fmt"
	"slices"

	"github.com/dizzrt/dauth/api/gen/errdef"
	"github.com/dizzrt/dauth/api/gen/token"
	"github.com/dizzrt/dauth/internal/infra/rpc/dauth"
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

func r(err error) {
	if err == nil {
		return
	}

	fmt.Printf("XAX: %s\n", err.Error())
	r(errors.Unwrap(err))
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

		err = errors.Unmarshal(err)
		r(err)

		if err != nil || resp.GetToken().Uid == 0 {
			// if err != nil && !(errors.Is(err, errdef.TokenExpired()) || errors.Is(err, errdef.TokenRevoked()) || errors.Is(err, jwt.ErrTokenExpired)) {
			// 	// BUG errors.Is not work as expected
			// 	log.CtxErrorf(ctx, "validate token failed, token: %v, err: %v", tokenStr, err)
			// }

			if errors.Is(err, errdef.TokenExpired()) {
				log.CtxInfo(ctx, "Hahaha")
			}

			unauthorized(ctx)
			return
		}

		ctx.Set("uid", resp.GetToken().Uid)
		ctx.Next()
	}
}
