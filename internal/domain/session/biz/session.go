package biz

import "github.com/dizzrt/dauth/internal/domain/session/repo"

var _ SessionBiz = (*sessionBiz)(nil)

type SessionBiz interface{}

type sessionBiz struct {
	sessionRepo repo.SessionRepo
}

func NewSessionBiz(sessionRepo repo.SessionRepo) SessionBiz {
	return &sessionBiz{
		sessionRepo: sessionRepo,
	}
}
