package entity

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          uint
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type UserRoleAssociation struct {
	ID        uint
	UserID    uint
	RoleID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
