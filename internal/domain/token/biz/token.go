package biz

import (
	"context"
	"time"
)

var _ TokenBiz = (*tokenBiz)(nil)

type TokenBiz interface {
	Issue(ctx context.Context, uid uint32, expiration time.Time) (string, error)
	Validate(ctx context.Context, token string, clientID string) (bool, string, error)
	Revoke(ctx context.Context, token string, reason string) (bool, error)
}

type tokenBiz struct {
}

func NewTokenBiz() TokenBiz {
	return &tokenBiz{}
}

func (biz *tokenBiz) Issue(ctx context.Context, uid uint32, expiration time.Time) (string, error) {
	return "", nil
}

func (biz *tokenBiz) Validate(ctx context.Context, token string, clientID string) (bool, string, error) {
	return false, "", nil
}

func (biz *tokenBiz) Revoke(ctx context.Context, token string, reason string) (bool, error) {
	return false, nil
}
