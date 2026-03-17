package cache

import (
	"context"
	"time"
)

type TokenCache interface {
	Revoke(ctx context.Context, token string, reason string, expiresAt time.Time) error
	IsRevoked(ctx context.Context, token string) (bool, string, error)
}
