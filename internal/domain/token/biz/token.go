package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/dizzrt/dauth/api/gen/errdef"
	token_api "github.com/dizzrt/dauth/api/gen/token"
	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/dauth/internal/domain/token/cache"
	"github.com/dizzrt/dauth/internal/domain/token/dto"
	"github.com/dizzrt/dauth/internal/domain/token/entity"
	"github.com/dizzrt/dauth/internal/domain/token/repo"
	"github.com/dizzrt/dauth/internal/infra/utils/security/jwt"
	"github.com/dizzrt/ellie/errors"
	"github.com/dizzrt/ellie/log"
	"github.com/google/uuid"
)

var _ TokenBiz = (*tokenBiz)(nil)
var tokenTTL time.Duration

func init() {
	ac := conf.GetAppConfig()

	var err error
	tokenTTL, err = time.ParseDuration(ac.App.TokenTTL)
	if err != nil {
		panic(err)
	}
}

type TokenBiz interface {
	IssueSSOToken(ctx context.Context, uid uint32) (token string, expiresAt time.Time, err error)
	Issue(ctx context.Context, uid uint32, clientID uint32, scope string) (accessToken string, refreshToken string, accessExpireAt, refreshExpireAt time.Time, err error)
	Validate(ctx context.Context, req *dto.ValidateRequest) (*entity.BaseToken, error)
	Revoke(ctx context.Context, token string, reason string) error
}

type tokenBiz struct {
	tokenBlacklistRepo repo.TokenBlacklistRepo
	tokenRevokeCache   cache.TokenRevokeCache
	jwtManager         jwt.JWTManager
}

func NewTokenBiz(tokenBlacklistRepo repo.TokenBlacklistRepo, tokenRevokeCache cache.TokenRevokeCache, jwtManager jwt.JWTManager) TokenBiz {
	return &tokenBiz{
		tokenBlacklistRepo: tokenBlacklistRepo,
		tokenRevokeCache:   tokenRevokeCache,
		jwtManager:         jwtManager,
	}
}

func (biz *tokenBiz) IssueSSOToken(ctx context.Context, uid uint32) (token string, expiresAt time.Time, err error) {
	now := time.Now()
	expiresAt = now.Add(tokenTTL)

	ssoToken := &entity.SSOToken{
		BaseToken: entity.BaseToken{
			RegisteredClaims: jwt.RegisteredClaims{
				ID:        uuid.NewString(),
				Issuer:    "dauth",
				Subject:   "dauth-sso",
				Audience:  jwt.ClaimStrings{"dauth"},
				IssuedAt:  jwt.NewNumericDate(now),
				NotBefore: jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(expiresAt),
			},
			UID:  uid,
			Type: token_api.Token_TokenType_SSO,
		},
	}

	ssoTokenStr, err := biz.jwtManager.Sign(ctx, ssoToken, nil)
	if err != nil {
		log.CtxErrorf(ctx, "sign sso token failed: %v", err)
		return
	}

	return ssoTokenStr, expiresAt, nil
}

func (biz *tokenBiz) Issue(ctx context.Context, uid uint32, clientID uint32, scope string) (accessToken string, refreshToken string, accessExpireAt, refreshExpireAt time.Time, err error) {
	// TODO read from config
	// accessExpire := 24 * time.Hour
	// refreshExpire := 7 * 24 * time.Hour

	// // validate service provider
	// resp, err := dauth.ValidateServiceProvider(ctx, uint32(clientID), scope)
	// if err != nil || !resp.GetIsOk() {
	// 	// invalid service provider or scope
	// 	return
	// }

	// now := time.Now()
	// accessExpireAt = now.Add(accessExpire)
	// accessTokenEntity := entity.Token{
	// 	TokenID:     uuid.NewString(),
	// 	UID:         uid,
	// 	ClientID:    clientID,
	// 	Issuer:      "dauth",
	// 	IssuedAt:    now,
	// 	NotBefore:   now,
	// 	ExpiresAt:   accessExpireAt,
	// 	Scope:       scope,
	// 	TokenType:   "Bearer",
	// 	Refreshable: false,
	// }

	// refreshExpireAt = now.Add(refreshExpire)
	// refreshTokenEntity := entity.Token{
	// 	TokenID:     uuid.NewString(),
	// 	UID:         uid,
	// 	ClientID:    clientID,
	// 	Issuer:      "dauth",
	// 	IssuedAt:    now,
	// 	NotBefore:   now,
	// 	ExpiresAt:   refreshExpireAt,
	// 	Scope:       scope,
	// 	TokenType:   "Bearer",
	// 	Refreshable: true,
	// }

	// accessToken, err = biz.jwtManager.Sign(ctx, accessTokenEntity.Claims(), nil)
	// if err != nil {
	// 	log.CtxErrorf(ctx, "sign access token failed: %v", err)
	// 	return
	// }

	// refreshToken, err = biz.jwtManager.Sign(ctx, refreshTokenEntity.Claims(), nil)
	// if err != nil {
	// 	log.CtxErrorf(ctx, "sign refresh token failed: %v", err)
	// 	return
	// }

	return
}

func (biz *tokenBiz) Validate(ctx context.Context, req *dto.ValidateRequest) (*entity.BaseToken, error) {
	var claims jwt.Claims
	switch req.TokenType {
	case token_api.Token_TokenType_SSO:
		claims = &entity.SSOToken{}
	// case token_api.Token_TokenType_ID:
	// 	claims = &entity.IDToken{}
	// case token_api.Token_TokenType_ACCESS:
	// 	claims = &entity.AccessToken{}
	// case token_api.Token_TokenType_REFRESH:
	// 	claims = &entity.RefreshToken{}
	default:
		return nil, fmt.Errorf("invalid token type: %v", req.TokenType)
	}

	err := biz.jwtManager.Verify(ctx, req.Token, nil, claims)
	if err != nil {
		return nil, err
	}

	// extract baseToken
	var baseToken *entity.BaseToken
	switch req.TokenType {
	case token_api.Token_TokenType_SSO:
		if ssoToken, ok := claims.(*entity.SSOToken); ok {
			baseToken = &ssoToken.BaseToken
		}
	// case token_api.Token_TokenType_ID:
	// 	if idToken, ok := claims.(*entity.IDToken); ok {
	// 		baseToken = &idToken.BaseToken
	// 	}
	// case token_api.Token_TokenType_ACCESS:
	// 	if accessToken, ok := claims.(*entity.AccessToken); ok {
	// 		baseToken = &accessToken.BaseToken
	// 	}
	// case token_api.Token_TokenType_REFRESH:
	// 	if refreshToken, ok := claims.(*entity.RefreshToken); ok {
	// 		baseToken = &refreshToken.BaseToken
	// 	}
	default:
		baseToken = nil
	}

	if baseToken == nil || baseToken.Type != req.TokenType {
		return nil, errdef.TokenInvalidWithMsg("token type not match")
	}

	isRevoked, _, err := biz.tokenRevokeCache.IsRevoked(ctx, req.Token)
	if err != nil || isRevoked {
		if err != nil {
			log.CtxErrorf(ctx, "check token if revoked failed, token: %v, err: %v", req.Token, err)
			return nil, err
		}

		// return nil, errdef.TokenRevoked()
		return nil, errdef.TokenRevoked()
	}

	return baseToken, nil
}

func (biz *tokenBiz) Revoke(ctx context.Context, token string, reason string) error {
	var claims entity.BaseToken
	err := biz.jwtManager.Verify(ctx, token, nil, &claims)
	if err != nil {
		if errors.Is(err, errdef.TokenExpired()) || errors.Is(err, errdef.TokenRevoked()) {
			return nil
		}

		return err
	}

	if err = biz.tokenRevokeCache.Revoke(ctx, token, reason, claims.ExpiresAt.Time); err != nil {
		log.CtxErrorf(ctx, "revoke token failed, token: %s, err: %v", token, err)
		return err
	}

	return nil
}
