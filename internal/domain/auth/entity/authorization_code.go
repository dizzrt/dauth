package entity

import (
	"encoding/json"
	"time"
)

type AuthorizationCode struct {
	Code        string
	UserID      uint32
	ClientID    uint32
	RedirectURI string
	Scope       string
	IssuedAt    time.Time
	ExpiresAt   time.Time
	Used        bool
}

func (ac *AuthorizationCode) MarshalBinary() ([]byte, error) {
	return json.Marshal(ac)
}

func (ac *AuthorizationCode) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, ac)
}
