package biz

import (
	"context"
	"time"

	"github.com/dizzrt/dauth/internal/domain/auth/cache"
	"github.com/dizzrt/dauth/internal/domain/auth/entity"
	"github.com/dizzrt/dauth/internal/domain/auth/repo"
	"github.com/dizzrt/ellie/log"
	"github.com/google/uuid"
)

var _ AuthBiz = (*authBiz)(nil)

type AuthBiz interface {
	GenerateAuthorizationCode(ctx context.Context, userID uint32, clientID uint32, redirectURI string, scope string) (string, error)
	ExchangeToken(ctx context.Context, code string, clientID uint32, clientSecret string, redirectURI string, grantType string) (accessToken, refreshToken string, accessExpireAt, refreshExpireAt time.Time, err error)
}

type authBiz struct {
	authorizationCodeRepo  repo.AuthorizationCodeRepo
	authorizationCodeCache cache.AuthorizationCodeCache
}

func NewAuthBiz(authCodeRepo repo.AuthorizationCodeRepo, authCodeCache cache.AuthorizationCodeCache) AuthBiz {
	return &authBiz{
		authorizationCodeRepo:  authCodeRepo,
		authorizationCodeCache: authCodeCache,
	}
}

func (biz *authBiz) GenerateAuthorizationCode(ctx context.Context, userID uint32, clientID uint32, redirectURI string, scope string) (string, error) {
	code := uuid.NewString()

	now := time.Now()
	authCode := &entity.AuthorizationCode{
		Code:        code,
		UserID:      userID,
		ClientID:    clientID,
		RedirectURI: redirectURI,
		Scope:       scope,
		IssuedAt:    now,
		ExpiresAt:   now.Add(time.Minute * 5),
		Used:        false,
	}

	err := biz.authorizationCodeCache.Set(ctx, code, authCode, time.Minute*5)
	if err != nil {
		log.CtxErrorf(ctx, "[GenerateAuthorizationCode] failed to set cache, err: %v", err)
		return "", err
	}

	return code, nil
}

func (biz *authBiz) ExchangeToken(ctx context.Context, code string, clientID uint32, clientSecret string, redirectURI string, grantType string) (accessToken, refreshToken string, accessExpireAt, refreshExpireAt time.Time, err error) {
	return "", "", time.Time{}, time.Time{}, nil
}
