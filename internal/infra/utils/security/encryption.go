package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const _BCRYPT_COST = 12

func GeneratePasswordHash(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("password can not be empty")
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(password), _BCRYPT_COST)
	if err != nil {
		return "", err
	}

	return string(pwd), nil
}
