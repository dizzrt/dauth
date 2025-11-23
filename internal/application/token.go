package application

import (
	"context"
	"fmt"

	"github.com/dizzrt/dauth/api/gen/token"
	"github.com/dizzrt/dauth/internal/domain/token/biz"
	"github.com/dizzrt/dauth/internal/infra/rpc"
)

var _ TokenApplication = (*tokenApplication)(nil)

type TokenApplication interface {
	Issue(context.Context, *token.IssueRequest) (*token.IssueResponse, error)
	Validate(context.Context, *token.ValidateRequest) (*token.ValidateResponse, error)
	Revoke(context.Context, *token.RevokeRequest) (*token.RevokeResponse, error)
}

type tokenApplication struct {
	tokenBiz biz.TokenBiz
}

func NewTokenApplication(tokenBiz biz.TokenBiz) TokenApplication {
	return &tokenApplication{
		tokenBiz: tokenBiz,
	}
}

func (app *tokenApplication) Issue(ctx context.Context, req *token.IssueRequest) (*token.IssueResponse, error) {
	if req.GetUid() == 0 || req.GetClientId() == 0 {
		return nil, fmt.Errorf("invalid params")
	}

	accessToken, refreshToken, accessExpireAt, refreshExpireAt, err := app.tokenBiz.Issue(ctx, req.GetUid(), req.GetClientId(), req.GetScope())
	if err != nil {
		return nil, err
	}

	return &token.IssueResponse{
		AccessToken:     accessToken,
		RefreshToken:    refreshToken,
		AccessExpireAt:  accessExpireAt.Unix(),
		RefreshExpireAt: refreshExpireAt.Unix(),
		BaseResp:        rpc.Success(),
	}, nil
}

func (app *tokenApplication) Validate(ctx context.Context, req *token.ValidateRequest) (*token.ValidateResponse, error) {
	ts := req.GetToken()
	clientID := req.GetClientId()

	if ts == "" || clientID == 0 {
		return nil, fmt.Errorf("invalid params")
	}

	tokenEntity, isValid, reason, err := app.tokenBiz.Validate(ctx, ts, req.GetClientId())
	if err != nil {
		return nil, err
	}

	return &token.ValidateResponse{
		Token: &token.Token{
			TokenId:     tokenEntity.TokenID,
			Uid:         tokenEntity.UID,
			ClientId:    tokenEntity.ClientID,
			Issuer:      tokenEntity.Issuer,
			IssuedAt:    tokenEntity.IssuedAt.Unix(),
			NotBefore:   tokenEntity.NotBefore.Unix(),
			ExpiresAt:   tokenEntity.ExpiresAt.Unix(),
			Scope:       tokenEntity.Scope,
			TokenType:   tokenEntity.TokenType,
			Refreshable: tokenEntity.Refreshable,
		},
		IsValid:  isValid,
		Reason:   reason,
		BaseResp: rpc.Success(),
	}, nil
}

func (app *tokenApplication) Revoke(ctx context.Context, req *token.RevokeRequest) (*token.RevokeResponse, error) {
	ts := req.GetToken()
	if ts == "" {
		return nil, fmt.Errorf("invalid params")
	}

	if err := app.tokenBiz.Revoke(ctx, ts, req.GetReason()); err != nil {
		return nil, err
	}

	return &token.RevokeResponse{
		IsSuccess: true,
		BaseResp:  rpc.Success(),
	}, nil
}
