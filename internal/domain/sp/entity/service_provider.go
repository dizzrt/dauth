package entity

import (
	"time"

	"github.com/dizzrt/dauth/api/gen/sp"
	"gorm.io/gorm"
)

type ServiceProvider struct {
	ID          uint32
	Name        string
	Description string
	Secret      string `validate:"required"`
	RedirectURI string
	Status      sp.ServiceProvider_Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
