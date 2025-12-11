package entity

import (
	"github.com/golang-jwt/jwt/v5"
)

var _ ClaimsableToken = (*SSOToken)(nil)

type SSOToken struct {
	TokenMeta
}

func NewSSOTokenFromClaims(claims jwt.Claims) (*SSOToken, error) {
	// ssoclaims, ok := claims.(*SSOClaims)
	// if !ok {
	// 	return nil, fmt.Errorf("claims is not *SSOClaims")
	// }

	return &SSOToken{}, nil
}

func (t *SSOToken) Claims() jwt.Claims {
	return jwt.RegisteredClaims{}

	// return &SSOClaims{
	// 	BaseClaims: BaseClaims{
	// 		Issuer:      t.Issuer,
	// 		IssuedAt:    t.IssuedAt,
	// 		ExpiresAt:   t.ExpiresAt,
	// 		NotBeforeAt: t.NotBeforeAt,
	// 		Subject:     t.Subject,
	// 		Audience:    t.Audience,
	// 	},
	// }
}
