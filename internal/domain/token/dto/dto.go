package dto

import (
	"github.com/dizzrt/dauth/api/gen/token"
)

type ValidateRequest struct {
	Token     string                `json:"token"`
	TokenType token.Token_TokenType `json:"type"`
	ClientID  uint32                `json:"client_id"`
}
