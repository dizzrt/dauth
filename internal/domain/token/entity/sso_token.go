package entity

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var _ jwt.Claims = (*SSOToken)(nil)

type SSOToken struct {
	BaseToken
}

func NewSSOTokenFromClaims(claims jwt.Claims) (*SSOToken, error) {
	token, ok := claims.(*SSOToken)
	if !ok {
		return nil, fmt.Errorf("claims is not *SSOToken")
	}

	if token.Type != TokenTypeSSO {
		return nil, fmt.Errorf("token type is not TokenTypeSSO")
	}

	return token, nil
}
