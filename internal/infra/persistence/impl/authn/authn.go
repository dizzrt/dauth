package authn

import (
	"github.com/dizzrt/dauth/internal/domain/authn/repo"
	"github.com/dizzrt/dauth/internal/infra/persistence/core"
)

var _ repo.AuthnRepo = (*AuthnRepoImpl)(nil)

type AuthnRepoImpl struct {
	core.RepoCore
}

func NewAuthnRepoImpl(repoCore core.RepoCore) repo.AuthnRepo {
	return &AuthnRepoImpl{
		RepoCore: repoCore,
	}
}
