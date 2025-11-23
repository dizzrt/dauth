package token

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ schema.Tabler = (*TokenBlacklist)(nil)

type TokenBlacklist struct {
	gorm.Model
	TokenID      string    `gorm:"column:token_id"`
	RevokedAt    time.Time `gorm:"column:revoked_at"`
	ExpiresAt    time.Time `gorm:"column:expires_at"`
	RevokeReason string    `gorm:"column:revoke_reason"`
}

func (TokenBlacklist) TableName() string {
	return "token_token_blacklist"
}
