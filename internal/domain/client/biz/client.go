package biz

import "github.com/dizzrt/dauth/internal/domain/client/repo"

var _ ClientBiz = (*clientBiz)(nil)

type ClientBiz interface{}

type clientBiz struct {
	clientRepo repo.ClientRepo
}

func NewClientBiz(clientRepo repo.ClientRepo) ClientBiz {
	return &clientBiz{
		clientRepo: clientRepo,
	}
}
