package client

import (
	"github.com/dizzrt/dauth/internal/domain/client/repo"
	"github.com/dizzrt/dauth/internal/infra/foundation"
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
