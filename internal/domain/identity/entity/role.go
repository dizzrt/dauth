package entity

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          uint32
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type UserRoleAssociation struct {
	ID        uint32
	UserID    uint32
	RoleID    uint32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
