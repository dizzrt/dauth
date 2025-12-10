package sp

import (
	"github.com/dizzrt/dauth/internal/domain/sp/entity"
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
	return "sp_scopes"
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
