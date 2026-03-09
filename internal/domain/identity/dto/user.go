package dto

import "github.com/dizzrt/dauth/internal/domain/identity/entity"

type CreateUserDTO struct {
	Password string
	Username string
	Nickname *string
	Phone    *string
	Email    *string
	Avatar   *string
}

type VerifyPasswordDTO struct {
	Password string
	UID      *uint32
	Username *string
	Email    *string
	Phone    *string
}

type VerifyPasswordResponse struct {
	OK                bool
	User              *entity.User
	IsLocked          *bool
	RemainingAttempts *int32
	RemainingLockTime *int64
}
