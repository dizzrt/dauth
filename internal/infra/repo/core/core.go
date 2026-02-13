package core

import (
	"github.com/dizzrt/dauth/internal/infra/repo/core/gen/dao"
	"gorm.io/gorm"
)

type RepoCore interface {
	Query() *dao.Query
}

type repoCore struct {
	q *dao.Query
}

func NewRepoCore(db *gorm.DB) RepoCore {
	return &repoCore{
		q: dao.Use(db),
	}
}

func (base *repoCore) Query() *dao.Query {
	return base.q
}
