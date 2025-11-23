package identity

import (
	"time"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ schema.Tabler = (*User)(nil)

type User struct {
	gorm.Model
	Email         string    `gorm:"unique"`
	Username      string    `gorm:"column:username"`
	Password      string    `gorm:"column:password"`
	Status        uint      `gorm:"column:status"`
	LastLoginTime time.Time `gorm:"column:last_login_time"`
}

func (u *User) TableName() string {
	return "identity_users"
}

// ToEntity converts the User model to the User entity.
func (u *User) ToEntity() *entity.User {
	return &entity.User{
		ID:            uint32(u.ID),
		Email:         u.Email,
		Username:      u.Username,
		Password:      u.Password,
		Status:        identity.User_Status(u.Status),
		LastLoginTime: u.LastLoginTime,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
		DeletedAt:     u.DeletedAt,
	}
}
