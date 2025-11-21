package entity

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ID          uint32
	ClientID    string `validate:"required"`
	Secret      string `validate:"required"`
	RedirectURI string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
