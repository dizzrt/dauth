package entity

import "time"

type AuthorizationCode struct {
	ID          uint32
	Code        string
	UserID      uint32
	ClientID    uint32
	RedirectURI string
	Scope       string
	IssuedAt    time.Time
	ExpiresAt   time.Time
	Used        bool
}
