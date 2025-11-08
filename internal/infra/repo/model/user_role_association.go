package model

import (
	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ schema.Tabler = (*UserRoleAssociation)(nil)

type UserRoleAssociation struct {
	gorm.Model
	UserID uint `gorm:"column:user_id"`
	RoleID uint `gorm:"column:role_id"`
}

func (u *UserRoleAssociation) TableName() string {
	return "identity_user_role_associations"
}

func (u *UserRoleAssociation) ToEntity() (*entity.UserRoleAssociation, error) {
	return &entity.UserRoleAssociation{
		ID:        u.ID,
		UserID:    u.UserID,
		RoleID:    u.RoleID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}, nil
}
