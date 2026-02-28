package identity

import (
	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"github.com/dizzrt/dauth/internal/infra/persistence/core/gen/model"
)

func fromIdentityUserEntity(entity *entity.User) *model.IdentityUser {
	return &model.IdentityUser{
		ID:          int32(entity.UID),
		Username:    entity.Username,
		Nickname:    entity.Nickname,
		Phone:       entity.Phone,
		Email:       entity.Email,
		Avatar:      entity.Avatar,
		Password:    entity.Password,
		Status:      int32(entity.Status),
		LastLoginAt: entity.LastLoginAt,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
	}
}

func toIdentityUserEntity(model *model.IdentityUser) *entity.User {
	return &entity.User{
		UID:         uint32(model.ID),
		Password:    model.Password,
		Username:    model.Username,
		Nickname:    model.Nickname,
		Phone:       model.Phone,
		Email:       model.Email,
		Avatar:      model.Avatar,
		Status:      identity.UserStatus(model.Status),
		LastLoginAt: model.LastLoginAt,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		DeletedAt:   model.DeletedAt,
	}
}

func toIdentityUserEntities(models []*model.IdentityUser) []*entity.User {
	entities := make([]*entity.User, 0, len(models))
	for _, model := range models {
		entities = append(entities, toIdentityUserEntity(model))
	}

	return entities
}
