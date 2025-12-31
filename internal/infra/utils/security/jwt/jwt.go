package jwt

import (
	"context"
	"errors"
	"fmt"

	"github.com/dizzrt/dauth/api/gen/errdef"
	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/dauth/internal/domain/token/cache"
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

	revokeCache cache.TokenRevokeCache
}

func NewJWTManager(ac *conf.AppConfig, revokeCache cache.TokenRevokeCache) JWTManager {
	// TODO read from config
	return &jwtManager{
		algorithm: _DEFAULT_ALGORITHM,
		issuer:    _DEFAULT_ISSUER,
		secret:    []byte(ac.App.Secret),
		// publicKey:  nil,
		// privateKey: nil,
		revokeCache: revokeCache,
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
		return errdef.TokenInvalid().WithMessage("claims is nil")
	}

	if secret == nil {
		secret = m.secret
	}

	jt, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != m.algorithm {
			return nil, errdef.TokenInvalid().WithMessage("unexpected signing method: %v", t.Header["alg"])
		}

		return secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return errdef.TokenExpired()
		}

		log.CtxErrorf(ctx, "parse token '%s' failed, err: %s", token, err.Error())
		return errdef.TokenInvalid().WithMessage("parse token failed").WithCause(err)
	}

	if !jt.Valid {
		return errdef.TokenExpired()
	}

	isRevoked, _, err := m.revokeCache.IsRevoked(ctx, token)
	if err != nil {
		log.CtxErrorf(ctx, "check token revoke cache failed: %s", err.Error())
		return err
	}

	if isRevoked {
		return errdef.TokenRevoked()
	}

	return nil
}
