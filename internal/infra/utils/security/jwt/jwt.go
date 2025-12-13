package jwt

import (
	"context"
	"fmt"

	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/ellie/log"
	"github.com/golang-jwt/jwt/v5"
)

var _ JWTManager = (*jwtManager)(nil)

const (
	_DEFAULT_ALGORITHM = "HS256"
	_DEFAULT_ISSUER    = "dauth"
)

type JWTManager interface {
	Sign(ctx context.Context, claims jwt.Claims, secret []byte) (string, error)
	Verify(ctx context.Context, token string, secret []byte, entity jwt.Claims) error
}

type jwtManager struct {
	algorithm string
	issuer    string
	secret    []byte // for HS256/HS384/HS512
	// publicKey  []byte // for RS256/RS512
	// privateKey []byte // for RS256/RS512
}

func NewJWTManager(ac *conf.AppConfig) JWTManager {
	// TODO read from config
	return &jwtManager{
		algorithm: _DEFAULT_ALGORITHM,
		issuer:    _DEFAULT_ISSUER,
		secret:    []byte(ac.App.Secret),
		// publicKey:  nil,
		// privateKey: nil,
	}
}

func (m *jwtManager) Sign(ctx context.Context, claims jwt.Claims, secret []byte) (string, error) {
	if secret == nil {
		secret = m.secret
	}

	signingMethod := jwt.GetSigningMethod(m.algorithm)
	if signingMethod == nil {
		log.CtxErrorf(ctx, "unknown algorithm: %s", m.algorithm)
		return "", fmt.Errorf("unknown algorithm: %s", m.algorithm)
	}

	token := jwt.NewWithClaims(signingMethod, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		log.CtxErrorf(ctx, "sign token failed with claims: %v; err: %v", claims, err)
		return "", err
	}

	return signedToken, nil
}

func (m *jwtManager) Verify(ctx context.Context, token string, secret []byte, claims jwt.Claims) error {
	if claims == nil {
		return fmt.Errorf("claims is nil")
	}

	if secret == nil {
		secret = m.secret
	}

	jt, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != m.algorithm {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return secret, nil
	})

	if err != nil {
		return fmt.Errorf("parse token failed: %w", err)
	}

	if !jt.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
