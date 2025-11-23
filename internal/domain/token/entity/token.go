package entity

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	TokenID     string // unique token id for field jti(jwt)
	UID         uint32
	ClientID    uint32
	Issuer      string
	IssuedAt    time.Time
	NotBefore   time.Time // invalid before this time
	ExpiresAt   time.Time
	Scope       string
	TokenType   string
	Refreshable bool
}

func (t *Token) Claims() jwt.Claims {
	return jwt.MapClaims{
		// standard claims
		"jti": t.TokenID,
		"sub": t.UID,
		"aud": t.ClientID,
		"iss": t.Issuer,
		"iat": t.IssuedAt.Unix(),
		"exp": t.ExpiresAt.Unix(),
		"nbf": t.NotBefore.Unix(),

		// custom claims
		"auth": map[string]any{
			"uid":    t.UID,
			"client": t.ClientID,
			"scope":  t.Scope,
		},
	}
}
