package session

import (
	"github.com/dizzrt/dauth/internal/domain/session/repo"
	"github.com/dizzrt/dauth/internal/infra/persistence/core"
)

var _ repo.SessionRepo = (*sessionRepoImpl)(nil)

type sessionRepoImpl struct {
	core.RepoCore
}

func NewSessionRepoImpl(repoCore core.RepoCore) repo.SessionRepo {
	return &sessionRepoImpl{
		RepoCore: repoCore,
	}
}
