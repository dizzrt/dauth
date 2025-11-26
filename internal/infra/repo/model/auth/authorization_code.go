package auth

import (
	"time"

	"github.com/dizzrt/dauth/internal/domain/auth/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ schema.Tabler = (*AuthorizationCode)(nil)

type AuthorizationCode struct {
	gorm.Model
	Code        string    `gorm:"column:code"`
	UserID      uint32    `gorm:"column:user_id"`
	ClientID    uint32    `gorm:"column:client_id"`
	RedirectURI string    `gorm:"column:redirect_uri"`
	Scope       string    `gorm:"column:scope"`
	IssuedAt    time.Time `gorm:"column:issued_at"`
	ExpiresAt   time.Time `gorm:"column:expires_at"`
	Used        bool      `gorm:"column:used"`
}

func (code *AuthorizationCode) TableName() string {
	return "auth_authorization_codes"
}

func (code *AuthorizationCode) ToEntity() *entity.AuthorizationCode {
	return &entity.AuthorizationCode{
		ID:          uint32(code.ID),
		Code:        code.Code,
		UserID:      code.UserID,
		ClientID:    code.ClientID,
		RedirectURI: code.RedirectURI,
		Scope:       code.Scope,
		IssuedAt:    code.IssuedAt,
		ExpiresAt:   code.ExpiresAt,
		Used:        code.Used,
	}
}
