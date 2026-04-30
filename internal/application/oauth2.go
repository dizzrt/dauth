package application

import "github.com/dizzrt/dauth/internal/domain/oauth2/biz"

var _ OAuth2Application = (*oauth2Application)(nil)

type OAuth2Application interface {
}

type oauth2Application struct {
	oauth2Biz biz.OAuth2Biz
}

func NewOAuth2Application(oauth2Biz biz.OAuth2Biz) OAuth2Application {
	return &oauth2Application{
		oauth2Biz: oauth2Biz,
	}
}
