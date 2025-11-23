package repo

import (
	"context"
	"time"
)

type TokenBlacklistRepo interface {
	Revoke(ctx context.Context, tokenID string, reason string, expiresAt time.Time) error
	IsRevoked(ctx context.Context, tokenID string) (bool, error)
}
