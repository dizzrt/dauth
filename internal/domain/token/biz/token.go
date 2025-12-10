package biz

import (
	"context"
	"errors"
	"time"

	"github.com/dizzrt/dauth/internal/domain/token/entity"
	"github.com/dizzrt/dauth/internal/domain/token/repo"
	"github.com/dizzrt/dauth/internal/infra/rpc/dauth"
	"github.com/dizzrt/dauth/internal/infra/utils/security/jwt"
	"github.com/dizzrt/ellie/log"
	"github.com/google/uuid"
)

var _ TokenBiz = (*tokenBiz)(nil)

type TokenBiz interface {
	Issue(ctx context.Context, uid uint32, clientID uint32, scope string) (accessToken string, refreshToken string, accessExpireAt, refreshExpireAt time.Time, err error)
	Validate(ctx context.Context, token string, clientID uint32) (*entity.Token, bool, string, error)
	Revoke(ctx context.Context, token string, reason string) error
}

type tokenBiz struct {
	tokenBlacklistRepo repo.TokenBlacklistRepo
	jwtManager         jwt.JWTManager
}

func NewTokenBiz(tokenBlacklistRepo repo.TokenBlacklistRepo, jwtManager jwt.JWTManager) TokenBiz {
	return &tokenBiz{
		tokenBlacklistRepo: tokenBlacklistRepo,
		jwtManager:         jwtManager,
	}
}

func (biz *tokenBiz) Issue(ctx context.Context, uid uint32, clientID uint32, scope string) (accessToken string, refreshToken string, accessExpireAt, refreshExpireAt time.Time, err error) {
	// TODO read from config
	accessExpire := 24 * time.Hour
	refreshExpire := 7 * 24 * time.Hour

	// validate service provider
	resp, err := dauth.ValidateServiceProvider(ctx, uint32(clientID), scope)
	if err != nil || !resp.GetIsOk() {
		// invalid service provider or scope
		return
	}

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

func (biz *tokenBiz) Validate(ctx context.Context, token string, clientID uint32) (*entity.Token, bool, string, error) {
	claims, err := biz.jwtManager.Verify(ctx, token)
	if err != nil {
		return nil, false, err.Error(), err
	}

	// convert claims to token entity
	tokenEntity, err := entity.NewTokenFromClaims(claims)
	if err != nil {
		log.CtxErrorf(ctx, "convert claims to token entity failed: %v", err)
		return nil, false, err.Error(), err
	}

	if tokenEntity.ClientID != clientID {
		return nil, false, "client id not match", errors.New("client id not match")
	}

	// TODO add cache check
	isRevoked, err := biz.tokenBlacklistRepo.IsRevoked(ctx, tokenEntity.TokenID)
	if err != nil || isRevoked {
		if err != nil {
			log.CtxErrorf(ctx, "check token blacklist failed, token: %v, err: %v", token, err)
		}

		return nil, false, "token is revoked", errors.New("token is revoked")
	}

	return tokenEntity, true, "", nil
}

func (biz *tokenBiz) Revoke(ctx context.Context, token string, reason string) error {
	claims, err := biz.jwtManager.Verify(ctx, token)
	if err != nil {
		log.CtxErrorf(ctx, "verify token failed: %v", err)
		return err
	}

	// convert claims to token entity
	tokenEntity, err := entity.NewTokenFromClaims(claims)
	if err != nil {
		log.CtxErrorf(ctx, "convert claims to token entity failed: %v", err)
		return err
	}

	tid := tokenEntity.TokenID
	if err := biz.tokenBlacklistRepo.Revoke(ctx, tid, reason, tokenEntity.ExpiresAt); err != nil {
		log.CtxErrorf(ctx, "revoke token failed, token: %v, err: %v", token, err)
		return err
	}

	return nil
}
