package jwt

import (
	"context"
	"fmt"

	"github.com/dizzrt/ellie/log"
	"github.com/golang-jwt/jwt/v5"
)

var _ JWTManager = (*jwtManager)(nil)

const (
	_DEFAULT_ALGORITHM = "HS256"
	_DEFAULT_ISSUER    = "dauth"
)

type JWTManager interface {
	Sign(ctx context.Context, claims jwt.Claims) (string, error)
	Verify(ctx context.Context, token string) (jwt.MapClaims, error)
}

type jwtManager struct {
	algorithm string
	issuer    string
	secret    []byte // for HS256/HS384/HS512
	// publicKey  []byte // for RS256/RS512
	// privateKey []byte // for RS256/RS512
}

func NewJWTManager() JWTManager {
	// TODO read from config
	return &jwtManager{
		algorithm: _DEFAULT_ALGORITHM,
		issuer:    _DEFAULT_ISSUER,
		secret:    []byte("hJB1u5iM7yj0DzdhL0YfRuDDY4BVwLRWcivrdaVvHeDPwkKbmQLcvp90F171"),
		// publicKey:  nil,
		// privateKey: nil,
	}
}

// Sign signs the claims with the given algorithm and secret.
func (m *jwtManager) Sign(ctx context.Context, claims jwt.Claims) (string, error) {
	signingMethod := jwt.GetSigningMethod(m.algorithm)
	if signingMethod == nil {
		log.CtxErrorf(ctx, "invalid algorithm: %s", m.algorithm)
		return "", fmt.Errorf("invalid algorithm: %s", m.algorithm)
	}

	token := jwt.NewWithClaims(signingMethod, claims)
	if m.issuer != "" {
		if mapClaims, ok := claims.(jwt.MapClaims); ok {
			mapClaims["iss"] = m.issuer
		}
	}

	signedToken, err := token.SignedString(m.secret)
	if err != nil {
		log.CtxErrorf(ctx, "sign token failed with claims: %v; err: %v", claims, err)
		return "", err
	}

	return signedToken, nil
}

func (m *jwtManager) Verify(ctx context.Context, token string) (jwt.MapClaims, error) {
	jt, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != m.algorithm {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return m.secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("parse token failed: %w", err)
	}

	if !jt.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := jt.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	// TODO check blacklist

	return claims, nil
}
