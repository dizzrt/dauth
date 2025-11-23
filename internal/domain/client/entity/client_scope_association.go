package entity

import (
	"time"

	"gorm.io/gorm"
)

type ClientScopeAssociation struct {
	ID        uint32
	ClientID  uint32
	ScopeID   uint32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
