package application

import (
	"context"
	"fmt"

	"github.com/dizzrt/dauth/api/gen/errdef"
	"github.com/dizzrt/dauth/api/gen/token"
	"github.com/dizzrt/dauth/internal/domain/token/biz"
	"github.com/dizzrt/dauth/internal/domain/token/dto"
	"github.com/dizzrt/dauth/internal/infra/rpc"
	"github.com/dizzrt/ellie/errors"
	"google.golang.org/grpc/codes"
)

var _ TokenApplication = (*tokenApplication)(nil)

type TokenApplication interface {
	IssueSSOToken(context.Context, *token.IssueSSOTokenRequest) (*token.IssueSSOTokenResponse, error)
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

func (app *tokenApplication) IssueSSOToken(ctx context.Context, req *token.IssueSSOTokenRequest) (*token.IssueSSOTokenResponse, error) {
	if req.GetUid() == 0 {
		return nil, errdef.InvalidArgument().WithMessage("uid is required")
	}

	tokenStr, expiresAt, err := app.tokenBiz.IssueSSOToken(ctx, req.GetUid())
	if err != nil {
		return nil, err
	}

	return &token.IssueSSOTokenResponse{
		Token:     tokenStr,
		ExpiresAt: expiresAt.Unix(),
		BaseResp:  rpc.Success(),
	}, nil
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
	tokenType := req.GetType()
	clientID := req.GetClientId()

	if ts == "" || tokenType == token.Token_TokenType_UNKNOWN {
		return nil, errdef.InvalidArgument()
	}

	if clientID == 0 && (tokenType == token.Token_TokenType_ID || tokenType == token.Token_TokenType_ACCESS || tokenType == token.Token_TokenType_REFRESH) {
		return nil, errdef.InvalidArgument().WithMessage("client_id is required when token_type is ID, ACCESS or REFRESH")
	}

	bt, err := app.tokenBiz.Validate(ctx, &dto.ValidateRequest{
		Token:     ts,
		ClientID:  clientID,
		TokenType: tokenType,
	})

	if err != nil {
		// return nil, err
		return nil, errors.Marshal(codes.Unauthenticated, err)
	}

	return &token.ValidateResponse{
		Token: &token.Token{
			Id:        bt.ID,
			Issuer:    bt.Issuer,
			Subject:   bt.Subject,
			Audience:  bt.Audience,
			IssuedAt:  bt.IssuedAt.Unix(),
			NotBefore: bt.NotBefore.Unix(),
			ExpiresAt: bt.ExpiresAt.Unix(),
			Uid:       bt.UID,
			Type:      bt.Type,
		},
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
