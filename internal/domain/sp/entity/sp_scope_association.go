package entity

import (
	"time"

	"gorm.io/gorm"
)

type SPScopeAssociation struct {
	ID        uint32
	SPID      uint32
	ScopeID   uint32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
