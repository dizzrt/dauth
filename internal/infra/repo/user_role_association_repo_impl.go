package repo

import (
	"github.com/dizzrt/dauth/internal/domain/identity/repo"
	"gorm.io/gorm"
)

var _ repo.UserRoleAssociationRepo = (*UserRoleAssociationRepoImpl)(nil)

type UserRoleAssociationRepoImpl struct {
	db *gorm.DB
}

func NewUserRoleAssociationRepoImpl(db *gorm.DB) repo.UserRoleAssociationRepo {
	return &UserRoleAssociationRepoImpl{db: db}
}
