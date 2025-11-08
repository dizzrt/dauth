package entity

import (
	"time"

	"github.com/dizzrt/dauth/api/gen/identity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID            uint32
	Email         string              `validate:"required,email"`
	Username      string              `validate:"required,max=20"`
	Password      string              `validate:"required"`
	Status        identity.UserStatus `validate:"required"`
	LastLoginTime time.Time
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
