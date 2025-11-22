package application

import "github.com/dizzrt/dauth/internal/domain/client/biz"

var _ ClientApplication = (*clientApplication)(nil)

type ClientApplication interface{}

type clientApplication struct {
	clientBiz biz.ClientBiz
}

func NewClientApplication(clientBiz biz.ClientBiz) ClientApplication {
	return &clientApplication{
		clientBiz: clientBiz,
	}
}
