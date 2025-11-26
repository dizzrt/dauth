package auth

import (
	"github.com/dizzrt/dauth/internal/domain/auth/repo"
	"github.com/dizzrt/dauth/internal/infra/foundation"
)

var _ repo.AuthorizationCodeRepo = (*AuthorizationCodeRepoImpl)(nil)

type AuthorizationCodeRepoImpl struct {
	*foundation.BaseDB
}

func NewAuthorizationCodeRepoImpl(base *foundation.BaseDB) repo.AuthorizationCodeRepo {
	return &AuthorizationCodeRepoImpl{
		BaseDB: base,
	}
}
