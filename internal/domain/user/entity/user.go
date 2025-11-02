package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint
	Email         string
	Username      string
	Password      string
	Status        uint
	LastLoginTime time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
