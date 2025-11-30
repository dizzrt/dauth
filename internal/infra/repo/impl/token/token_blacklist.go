package token

import (
	"context"
	"errors"
	"time"

	"github.com/dizzrt/dauth/internal/domain/token/repo"
	"github.com/dizzrt/dauth/internal/infra/foundation"
	model "github.com/dizzrt/dauth/internal/infra/repo/model/token"
	"gorm.io/gorm"
)

var _ repo.TokenBlacklistRepo = (*TokenBlacklistImpl)(nil)

type TokenBlacklistImpl struct {
	*foundation.BaseDB
}

func NewTokenBlacklistRepoImpl(base *foundation.BaseDB) repo.TokenBlacklistRepo {
	return &TokenBlacklistImpl{
		BaseDB: base,
	}
}

func (impl *TokenBlacklistImpl) Revoke(ctx context.Context, tokenID string, reason string, expiresAt time.Time) error {
	m := &model.TokenBlacklist{
		TokenID:      tokenID,
		RevokedAt:    time.Now(),
		ExpiresAt:    expiresAt,
		RevokeReason: reason,
	}

	db := impl.WithContext(ctx)
	if err := db.Create(&m).Error; err != nil {
		return impl.WrapError(err)
	}

	return nil
}

func (impl *TokenBlacklistImpl) IsRevoked(ctx context.Context, tokenID string) (bool, error) {
	var m *model.TokenBlacklist

	db := impl.WithContext(ctx)
	if err := db.Where("token_id = ?", tokenID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, impl.WrapError(err)
	}

	return true, nil
}
