package biz

import (
	"context"
	"time"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/domain/authn/dto"
	"github.com/dizzrt/dauth/internal/domain/authn/repo"
	"github.com/dizzrt/dauth/internal/infra/rpc/dauth"
	"github.com/dizzrt/dauth/internal/infra/utils/security/jwt"
	"github.com/dizzrt/ellie/log"
	"github.com/rs/xid"
)

var _ AuthnBiz = (*authnBiz)(nil)

type AuthnBiz interface {
	Login(ctx context.Context, account string, password string) (*dto.LoginResponse, error)
	CheckAuthnStatus(ctx context.Context, token string) error
}

type authnBiz struct {
	jwtManager jwt.JWTManager

	authnRepo repo.AuthnRepo
}

func NewAuthnBiz(jwtManager jwt.JWTManager, authnRepo repo.AuthnRepo) AuthnBiz {
	return &authnBiz{jwtManager: jwtManager, authnRepo: authnRepo}
}

func (biz *authnBiz) Login(ctx context.Context, account string, password string) (*dto.LoginResponse, error) {
	loginResp := &dto.LoginResponse{
		Status: dto.LoginStatusFailed,
	}

	// TODO read authn policy

	resp, err := dauth.VerifyPassword(ctx, &identity.VerifyPasswordRequest{
		Password: password,
		Username: &account, // only support username login yet
	})

	if err != nil {
		log.CtxErrorf(ctx, "[Authn] verify password failed, account: %s, err: %v", account, err)
		return loginResp, err
	}

	if !resp.GetOk() {
		// TODO handle lockout
		return loginResp, nil
	}

	// TODO create session
	sessionID := xid.New().String()

	now := time.Now()
	expiresAt := now.Add(24 * time.Hour)
	idToken := jwt.IDToken{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        xid.New().String(),
			Issuer:    "dauth",
			Subject:   "dauth",
			Audience:  jwt.ClaimStrings{"dauth"},
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		Type: jwt.TokenTypeID,
		UID:  resp.GetUser().GetUid(),
		SID:  sessionID,
	}

	token, err := biz.jwtManager.Sign(ctx, idToken, []byte(""))
	if err != nil {
		log.CtxErrorf(ctx, "[Authn] sign id token failed, account: %s, err: %v", account, err)
		return loginResp, err
	}

	loginResp.Token = token
	loginResp.Status = dto.LoginStatusSuccess
	return loginResp, nil
}

func (biz *authnBiz) CheckAuthnStatus(ctx context.Context, token string) error {
	idToken := &jwt.IDToken{}
	err := biz.jwtManager.Verify(ctx, token, jwt.TokenTypeID, idToken, nil)

	return err
}
