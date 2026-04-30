package oauth2

import (
	"github.com/dizzrt/dauth/internal/domain/oauth2/repo"
	"github.com/dizzrt/dauth/internal/infra/persistence/core"
)

var _ repo.OAuth2Repo = (*oauth2RepoImpl)(nil)

type oauth2RepoImpl struct {
	core.RepoCore
}

func NewOAuth2RepoImpl(repoCore core.RepoCore) repo.OAuth2Repo {
	return &oauth2RepoImpl{
		RepoCore: repoCore,
	}
}
