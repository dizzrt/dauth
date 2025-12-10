package sp

import (
	"context"

	"github.com/dizzrt/dauth/internal/domain/sp/entity"
	"github.com/dizzrt/dauth/internal/domain/sp/repo"
	"github.com/dizzrt/dauth/internal/infra/foundation"
	model "github.com/dizzrt/dauth/internal/infra/repo/model/sp"
)

var _ repo.SPScopeAssociationRepo = (*SPScopeAssociationRepoImpl)(nil)

type SPScopeAssociationRepoImpl struct {
	*foundation.BaseDB
}

func NewSPScopeAssociationRepoImpl(base *foundation.BaseDB) repo.SPScopeAssociationRepo {
	return &SPScopeAssociationRepoImpl{
		BaseDB: base,
	}
}

func (impl *SPScopeAssociationRepoImpl) CreateAssociations(ctx context.Context, spID uint32, scopeIDs []uint32) error {
	if len(scopeIDs) == 0 {
		return nil
	}

	scopes := make([]*model.SPScopeAssociation, 0, len(scopeIDs))
	for _, id := range scopeIDs {
		scopes = append(scopes, &model.SPScopeAssociation{
			SPID:    spID,
			ScopeID: id,
		})
	}

	db := impl.WithContext(ctx)
	if err := db.Create(&scopes).Error; err != nil {
		return impl.WrapError(err)
	}

	return nil
}

func (impl *SPScopeAssociationRepoImpl) GetSPScopes(ctx context.Context, spID uint32) ([]*entity.SPScopeAssociation, error) {
	var associations []*model.SPScopeAssociation

	db := impl.WithContext(ctx)
	if err := db.Where("sp_id = ?", spID).Find(&associations).Error; err != nil {
		return nil, impl.WrapError(err)
	}

	associationsEntity := make([]*entity.SPScopeAssociation, 0, len(associations))
	for _, a := range associations {
		associationsEntity = append(associationsEntity, a.ToEntity())
	}

	return associationsEntity, nil
}
