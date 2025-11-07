package repo

import (
	"github.com/dizzrt/dauth/internal/domain/identity/repo"
	"gorm.io/gorm"
)

var _ repo.RoleRepo = (*RoleRepoImpl)(nil)

type RoleRepoImpl struct {
	db *gorm.DB
}

func NewRoleRepoImpl(db *gorm.DB) repo.RoleRepo {
	return &RoleRepoImpl{db: db}
}
