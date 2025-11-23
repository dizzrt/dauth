package biz

import "github.com/dizzrt/dauth/internal/domain/client/repo"

var _ ClientBiz = (*clientBiz)(nil)

type ClientBiz interface{}

type clientBiz struct {
	clientRepo                 repo.ClientRepo
	scopeRepo                  repo.ScopeRepo
	clientScopeAssociationRepo repo.ClientScopeAssociationRepo
}

func NewClientBiz(clientRepo repo.ClientRepo, scopeRepo repo.ScopeRepo, clientScopeAssociationRepo repo.ClientScopeAssociationRepo) ClientBiz {
	return &clientBiz{
		clientRepo:                 clientRepo,
		scopeRepo:                  scopeRepo,
		clientScopeAssociationRepo: clientScopeAssociationRepo,
	}
}
