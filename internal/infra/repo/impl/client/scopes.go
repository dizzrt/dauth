package client

import (
	"context"

	"github.com/dizzrt/dauth/internal/domain/client/entity"
	"github.com/dizzrt/dauth/internal/domain/client/repo"
	"github.com/dizzrt/dauth/internal/infra/foundation"
	model "github.com/dizzrt/dauth/internal/infra/repo/model/client"
)

var _ repo.ScopeRepo = (*ScopeRepoImpl)(nil)

type ScopeRepoImpl struct {
	*foundation.BaseDB
}

func NewScopeRepoImpl(base *foundation.BaseDB) repo.ScopeRepo {
	return &ScopeRepoImpl{
		BaseDB: base,
	}
}

func (impl *ScopeRepoImpl) GetScopesByIDs(ctx context.Context, ids []uint32) ([]*entity.Scope, error) {
	var scopes []*model.Scope

	db := impl.WithContext(ctx)
	if err := db.Where("id IN ? AND deleted_at IS NULL", ids).Find(&scopes).Error; err != nil {
		return nil, err
	}

	scopesEntity := make([]*entity.Scope, 0, len(scopes))
	for _, s := range scopes {
		scopesEntity = append(scopesEntity, s.ToEntity())
	}

	return scopesEntity, nil
}
