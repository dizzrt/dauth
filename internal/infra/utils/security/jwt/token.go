package jwt

import "github.com/golang-jwt/jwt/v5"

type TokenType uint8

const (
	TokenType_ID TokenType = iota + 1
	TokenType_ACCESS
	TokenType_REFRESH
)

type Token struct {
	jwt.RegisteredClaims
	Type TokenType `json:"type"`
}

type IDToken struct {
	Token

	UID uint32 `json:"uid"`
	SID string `json:"sid"` // session id
}
