package entity

import (
	"time"

	"github.com/dizzrt/dauth/api/gen/identity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID            uint
	Email         string
	Username      string
	Password      string
	Status        identity.UserStatus
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
