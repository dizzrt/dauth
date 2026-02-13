package entity

import (
	"time"

	"github.com/dizzrt/dauth/api/gen/identity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UID      uint32 `validate:"required"`
	Password string `validate:"required"`

	Username *string `validate:"required,max=32"`
	Nickname *string `validate:"required,max=32"`
	Phone    *string `validate:"required,phone"`
	Email    *string `validate:"required,email"`
	Avatar   *string `validate:"required,url"`

	Status        identity.UserStatus `validate:"required"`
	LastLoginTime *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

func (u *User) VerifyPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}
