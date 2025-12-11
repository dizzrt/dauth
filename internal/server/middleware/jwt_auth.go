package middleware

import (
	"slices"

	"github.com/dizzrt/ellie/transport/http"
	"github.com/gin-gonic/gin"
)

var authWhiteList = []string{}

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

		//   token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		//      return []byte("secret"), nil
		//   })

		//   if err != nil {
		//      unauthorized(ctx)
		//      return
		//   }

		//   if claims, ok := token.Claims.(*JWTClaims); !ok || !token.Valid {
		//      unauthorized(ctx)
		//      return
		// }

		//	ctx.Set("id", claims.ID)
		//	ctx.Set("user_name", claims.UserName)
		ctx.Next()
	}
}
