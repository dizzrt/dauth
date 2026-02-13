package dto

type CreateUserDTO struct {
	Password string
	Username *string
	Nickname *string
	Phone    *string
	Email    *string
	Avatar   *string
}
