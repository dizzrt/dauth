package client

import (
	"github.com/dizzrt/dauth/internal/domain/client/repo"
	"github.com/dizzrt/dauth/internal/infra/foundation"
)

var _ repo.ClientRepo = (*ClientRepoImpl)(nil)

type ClientRepoImpl struct {
	*foundation.BaseDB
}

func NewClientRepoImpl(base *foundation.BaseDB) repo.ClientRepo {
	return &ClientRepoImpl{
		BaseDB: base,
	}
}
