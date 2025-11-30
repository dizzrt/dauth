package client

import (
	"context"

	"github.com/dizzrt/dauth/internal/domain/client/entity"
	"github.com/dizzrt/dauth/internal/domain/client/repo"
	"github.com/dizzrt/dauth/internal/infra/foundation"
	model "github.com/dizzrt/dauth/internal/infra/repo/model/client"
)

var _ repo.ClientScopeAssociationRepo = (*ClientScopeAssociationRepoImpl)(nil)

type ClientScopeAssociationRepoImpl struct {
	*foundation.BaseDB
}

func NewClientScopeAssociationRepoImpl(base *foundation.BaseDB) repo.ClientScopeAssociationRepo {
	return &ClientScopeAssociationRepoImpl{
		BaseDB: base,
	}
}

func (impl *ClientScopeAssociationRepoImpl) CreateAssociations(ctx context.Context, clientID uint32, scopeIDs []uint32) error {
	if len(scopeIDs) == 0 {
		return nil
	}

	scopes := make([]*model.ClientScopeAssociation, 0, len(scopeIDs))
	for _, id := range scopeIDs {
		scopes = append(scopes, &model.ClientScopeAssociation{
			ClientID: clientID,
			ScopeID:  id,
		})
	}

	db := impl.WithContext(ctx)
	if err := db.Create(&scopes).Error; err != nil {
		return impl.WrapError(err)
	}

	return nil
}

func (impl *ClientScopeAssociationRepoImpl) GetClientScopes(ctx context.Context, clientID uint32) ([]*entity.ClientScopeAssociation, error) {
	var associations []*model.ClientScopeAssociation

	db := impl.WithContext(ctx)
	if err := db.Where("client_id = ?", clientID).Find(&associations).Error; err != nil {
		return nil, impl.WrapError(err)
	}

	associationsEntity := make([]*entity.ClientScopeAssociation, 0, len(associations))
	for _, a := range associations {
		associationsEntity = append(associationsEntity, a.ToEntity())
	}

	return associationsEntity, nil
}
