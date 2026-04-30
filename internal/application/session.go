package application

import "github.com/dizzrt/dauth/internal/domain/session/biz"

var _ SessionApplication = (*sessionApplication)(nil)

type SessionApplication interface {
}

type sessionApplication struct {
	sessionBiz biz.SessionBiz
}

func NewSessionApplication(sessionBiz biz.SessionBiz) OAuth2Application {
	return &sessionApplication{
		sessionBiz: sessionBiz,
	}
}
