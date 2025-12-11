package entity

import (
	"github.com/golang-jwt/jwt/v5"
)

var _ jwt.Claims = (*SSOClaims)(nil)

type SSOClaims struct {
	jwt.RegisteredClaims
	SSO struct {
		UserID uint32 `json:"user_id"`
	} `json:"sso"`
}
