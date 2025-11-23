package client

import (
	"github.com/dizzrt/dauth/internal/domain/client/repo"
	"github.com/dizzrt/dauth/internal/infra/foundation"
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
