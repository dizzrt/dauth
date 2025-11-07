package biz

import "github.com/dizzrt/dauth/internal/domain/identity/repo"

type RoleBiz interface {
}

type roleBiz struct {
	roleRepo                repo.RoleRepo
	userRoleAssociationRepo repo.UserRoleAssociationRepo
}

func NewRoleBiz(roleRepo repo.RoleRepo, userRoleAssociationRepo repo.UserRoleAssociationRepo) RoleBiz {
	return &roleBiz{
		roleRepo:                roleRepo,
		userRoleAssociationRepo: userRoleAssociationRepo,
	}
}
