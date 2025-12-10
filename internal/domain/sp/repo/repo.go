package repo

import (
	"context"

	"github.com/dizzrt/dauth/internal/domain/sp/entity"
)

type ServiceProviderRepo interface {
	Create(ctx context.Context, sp *entity.ServiceProvider) (*entity.ServiceProvider, error)
	Get(ctx context.Context, id uint32) (*entity.ServiceProvider, error)
	List(ctx context.Context, size, offset uint32) ([]*entity.ServiceProvider, uint32, error)
}

type ScopeRepo interface {
	GetScopesByIDs(ctx context.Context, ids []uint32) ([]*entity.Scope, error)
}

type SPScopeAssociationRepo interface {
	CreateAssociations(ctx context.Context, spID uint32, scopeIDs []uint32) error
	GetSPScopes(ctx context.Context, spID uint32) ([]*entity.SPScopeAssociation, error)
}
