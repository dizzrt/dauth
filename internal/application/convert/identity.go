package convert

import (
	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/domain/identity/entity"
)

func RolesToIdentityRoles(roles []*entity.Role) []*identity.Role {
	if roles == nil {
		return nil
	}

	var identityRoles []*identity.Role
	for _, role := range roles {
		identityRoles = append(identityRoles, &identity.Role{
			Id:          role.ID,
			Name:        role.Name,
			Description: role.Description,
		})
	}

	return identityRoles
}

func UserToIdentityUser(user *entity.User) *identity.User {
	return &identity.User{
		Id:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		Status:      user.Status,
		LastLoginAt: user.LastLoginTime.Unix(),
		CreatedAt:   user.CreatedAt.Unix(),
		UpdatedAt:   user.UpdatedAt.Unix(),
	}
}
