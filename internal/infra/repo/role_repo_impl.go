package repo

import (
	"context"

	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"github.com/dizzrt/dauth/internal/domain/identity/repo"
	"github.com/dizzrt/dauth/internal/infra/common"
	"github.com/dizzrt/dauth/internal/infra/repo/model"
)

var _ repo.RoleRepo = (*RoleRepoImpl)(nil)

type RoleRepoImpl struct {
	*common.BaseDB
}

func NewRoleRepoImpl(db *common.BaseDB) repo.RoleRepo {
	return &RoleRepoImpl{
		BaseDB: db,
	}
}

func (impl *RoleRepoImpl) CreateRole(ctx context.Context, role *entity.Role) (uint, error) {
	m := &model.Role{
		Name:        role.Name,
		Description: role.Description,
	}

	db := impl.WithContext(ctx)
	if err := db.Create(m).Error; err != nil {
		return 0, err
	}

	return m.ID, nil
}

func (impl *RoleRepoImpl) ListRolesWithPage(ctx context.Context, page, pageSize int) ([]*entity.Role, error) {
	var ms []model.Role

	db := impl.WithContext(ctx)
	err := db.Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&ms).Error

	if err != nil {
		return nil, err
	}

	roles := make([]*entity.Role, 0, len(ms))
	for _, m := range ms {
		role, err := m.ToEntity()
		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}

func (impl *RoleRepoImpl) DeleteRoles(ctx context.Context, ids []uint) error {
	db := impl.WithContext(ctx)
	err := db.Delete(&model.Role{}, "id IN ?", ids).Error

	return err
}

func (impl *RoleRepoImpl) UpdateRole(ctx context.Context, role *entity.Role) error {
	db := impl.WithContext(ctx)
	err := db.Model(&model.Role{}).
		Where("id = ?", role.ID).
		Updates(map[string]any{
			"name":        role.Name,
			"description": role.Description,
		}).Error

	return err
}
