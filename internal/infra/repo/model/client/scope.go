package client

import (
	"github.com/dizzrt/dauth/internal/domain/client/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ schema.Tabler = (*Scope)(nil)

type Scope struct {
	gorm.Model
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}

func (s *Scope) TableName() string {
	return "client_scopes"
}

func (s *Scope) ToEntity() *entity.Scope {
	return &entity.Scope{
		ID:          uint32(s.ID),
		Name:        s.Name,
		Description: s.Description,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
		DeletedAt:   s.DeletedAt,
	}
}
