package entity

import (
	"time"

	"gorm.io/gorm"
)

type Scope struct {
	ID          uint32
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
