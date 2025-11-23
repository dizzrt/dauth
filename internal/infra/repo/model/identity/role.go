package identity

import (
	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ schema.Tabler = (*Role)(nil)

type Role struct {
	gorm.Model
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}

func (r *Role) TableName() string {
	return "identity_roles"
}

func (r *Role) ToEntity() (*entity.Role, error) {
	return &entity.Role{
		ID:          uint32(r.ID),
		Name:        r.Name,
		Description: r.Description,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		DeletedAt:   r.DeletedAt,
	}, nil
}
