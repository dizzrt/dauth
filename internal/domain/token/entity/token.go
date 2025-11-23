package entity

import (
	"encoding/json"
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

func NewTokenFromClaims(claims jwt.MapClaims) (*Token, error) {
	issuedAt, err := claims.GetIssuedAt()
	if err != nil {
		return nil, err
	}

	notBefore, err := claims.GetNotBefore()
	if err != nil {
		return nil, err
	}

	expiresAt, err := claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	jb, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}

	var mp map[string]any
	err = json.Unmarshal(jb, &mp)
	if err != nil {
		return nil, err
	}

	auth := mp["auth"].(map[string]any)
	token := &Token{
		TokenID:     mp["jti"].(string),
		UID:         uint32(auth["uid"].(float64)),
		ClientID:    uint32(auth["client"].(float64)),
		Issuer:      mp["iss"].(string),
		IssuedAt:    issuedAt.Time,
		NotBefore:   notBefore.Time,
		ExpiresAt:   expiresAt.Time,
		Scope:       auth["scope"].(string),
		TokenType:   auth["type"].(string),
		Refreshable: auth["refreshable"].(bool),
	}

	return token, nil
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
			"uid":         t.UID,
			"client":      t.ClientID,
			"scope":       t.Scope,
			"type":        t.TokenType,
			"refreshable": t.Refreshable,
		},
	}
}
