package jwt

import "github.com/golang-jwt/jwt/v5"

type TokenType uint8

const (
	TokenTypeID TokenType = iota + 1
	TokenTypeAccess
	TokenTypeRefresh
)

type IDToken struct {
	jwt.RegisteredClaims
	Type TokenType `json:"type"`

	UID uint32 `json:"uid"`
	SID string `json:"sid"` // session id
}
