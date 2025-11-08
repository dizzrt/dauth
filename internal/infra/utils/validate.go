package utils

import (
	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func init() {
	v = validator.New()
}

func Validate() *validator.Validate {
	return v
}
