package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type NumericDate = jwt.NumericDate
type ClaimStrings = jwt.ClaimStrings
type RegisteredClaims = jwt.RegisteredClaims

func NewNumericDate(t time.Time) *NumericDate {
	return jwt.NewNumericDate(t)
}
