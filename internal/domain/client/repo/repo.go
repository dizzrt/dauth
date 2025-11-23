package repo

import (
	"context"

	"github.com/dizzrt/dauth/internal/domain/client/entity"
)

type ClientRepo interface {
	Create(ctx context.Context, client *entity.Client) (*entity.Client, error)
	Get(ctx context.Context, id uint32) (*entity.Client, error)
}

type ScopeRepo interface {
	GetScopesByIDs(ctx context.Context, ids []uint32) ([]*entity.Scope, error)
}

type ClientScopeAssociationRepo interface {
	CreateAssociations(ctx context.Context, clientID uint32, scopeIDs []uint32) error
	GetClientScopes(ctx context.Context, clientID uint32) ([]*entity.ClientScopeAssociation, error)
}
