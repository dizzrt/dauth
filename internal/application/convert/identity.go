package convert

import (
	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToIdentityUser(user *entity.User) *identity.User {
	u := &identity.User{
		Uid:       &user.UID,
		Username:  &user.Username,
		Status:    &user.Status,
		Phone:     user.Phone,
		Email:     user.Email,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Extend:    nil,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}

	if user.LastLoginAt != nil {
		u.LastLoginAt = timestamppb.New(*user.LastLoginAt)
	}

	if user.DeletedAt.Valid {
		u.DeletedAt = timestamppb.New(user.DeletedAt.Time)
	}

	return u
}

func ToIdentityUsers(users []*entity.User) []*identity.User {
	us := make([]*identity.User, 0, len(users))
	for _, user := range users {
		us = append(us, ToIdentityUser(user))
	}

	return us
}
