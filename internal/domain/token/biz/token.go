package biz

import (
	"context"
	"time"

	"github.com/dizzrt/dauth/internal/domain/token/entity"
	"github.com/dizzrt/dauth/internal/infra/utils/security/jwt"
	"github.com/dizzrt/ellie/log"
	"github.com/google/uuid"
)

var _ TokenBiz = (*tokenBiz)(nil)

type TokenBiz interface {
	Issue(ctx context.Context, uid uint32, clientID string, scope string) (accessToken string, refreshToken string, accessExpireAt, refreshExpireAt time.Time, err error)
	Validate(ctx context.Context, token string, clientID string) (bool, string, error)
	Revoke(ctx context.Context, token string, reason string) (bool, error)
}

type tokenBiz struct {
	jwtManager jwt.JWTManager
}

func NewTokenBiz(jwtManager jwt.JWTManager) TokenBiz {
	return &tokenBiz{
		jwtManager: jwtManager,
	}
}

func (biz *tokenBiz) Issue(ctx context.Context, uid uint32, clientID string, scope string) (accessToken string, refreshToken string, accessExpireAt, refreshExpireAt time.Time, err error) {
	// TODO read from config
	accessExpire := 24 * time.Hour
	refreshExpire := 7 * 24 * time.Hour

	// TODO verify client

	// TODO verify scope

	now := time.Now()
	accessExpireAt = now.Add(accessExpire)
	accessTokenEntity := entity.Token{
		TokenID:     uuid.NewString(),
		UID:         uid,
		ClientID:    clientID,
		Issuer:      "dauth",
		IssuedAt:    now,
		NotBefore:   now,
		ExpiresAt:   accessExpireAt,
		Scope:       scope,
		TokenType:   "Bearer",
		Refreshable: false,
	}

	refreshExpireAt = now.Add(refreshExpire)
	refreshTokenEntity := entity.Token{
		TokenID:     uuid.NewString(),
		UID:         uid,
		ClientID:    clientID,
		Issuer:      "dauth",
		IssuedAt:    now,
		NotBefore:   now,
		ExpiresAt:   refreshExpireAt,
		Scope:       scope,
		TokenType:   "Bearer",
		Refreshable: true,
	}

	accessToken, err = biz.jwtManager.Sign(ctx, accessTokenEntity.Claims())
	if err != nil {
		log.CtxErrorf(ctx, "sign access token failed: %v", err)
		return
	}

	refreshToken, err = biz.jwtManager.Sign(ctx, refreshTokenEntity.Claims())
	if err != nil {
		log.CtxErrorf(ctx, "sign refresh token failed: %v", err)
		return
	}

	return
}

func (biz *tokenBiz) Validate(ctx context.Context, token string, clientID string) (bool, string, error) {
	return false, "", nil
}

func (biz *tokenBiz) Revoke(ctx context.Context, token string, reason string) (bool, error) {
	return false, nil
}
