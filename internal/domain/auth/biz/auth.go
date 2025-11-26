package biz

import (
	"context"
	"time"

	"github.com/dizzrt/dauth/internal/domain/auth/repo"
)

var _ AuthBiz = (*authBiz)(nil)

type AuthBiz interface {
	GenerateAuthorizationCode(ctx context.Context, userID uint32, clientID uint32, redirectURI string, scope string) (string, error)
	ExchangeToken(ctx context.Context, code string, clientID uint32, clientSecret string, redirectURI string, grantType string) (accessToken, refreshToken string, accessExpireAt, refreshExpireAt time.Time, err error)
}

type authBiz struct {
	authorizationCodeRepo repo.AuthorizationCodeRepo
}

func NewAuthBiz(authRepo repo.AuthorizationCodeRepo) AuthBiz {
	return &authBiz{
		authorizationCodeRepo: authRepo,
	}
}

func (biz *authBiz) GenerateAuthorizationCode(ctx context.Context, userID uint32, clientID uint32, redirectURI string, scope string) (string, error) {
	return "", nil
}

func (biz *authBiz) ExchangeToken(ctx context.Context, code string, clientID uint32, clientSecret string, redirectURI string, grantType string) (accessToken, refreshToken string, accessExpireAt, refreshExpireAt time.Time, err error) {
	return "", "", time.Time{}, time.Time{}, nil
}
