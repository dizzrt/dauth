package identity

import (
	"context"

	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"github.com/dizzrt/dauth/internal/domain/identity/repo"
	"github.com/dizzrt/dauth/internal/infra/foundation"
)

var _ repo.UserRoleAssociationRepo = (*UserRoleAssociationRepoImpl)(nil)

type UserRoleAssociationRepoImpl struct {
	*foundation.BaseDB
}

func NewUserRoleAssociationRepoImpl(base *foundation.BaseDB) repo.UserRoleAssociationRepo {
	return &UserRoleAssociationRepoImpl{
		BaseDB: base,
	}
}

func (impl *UserRoleAssociationRepoImpl) AssignRoles(ctx context.Context, uid uint, roleIDs []uint) error {
	// TODO: Implement AssignRoles method.

	return nil
}

func (impl *UserRoleAssociationRepoImpl) UnassignRoles(ctx context.Context, uid uint, roleIDs []uint) error {
	// TODO: Implement UnassignRoles method.

	return nil
}

func (impl *UserRoleAssociationRepoImpl) GetUserRoles(ctx context.Context, uid uint) ([]*entity.Role, error) {
	// TODO: Implement GetUserRoles method.

	return nil, nil
}
