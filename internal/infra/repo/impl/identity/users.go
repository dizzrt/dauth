package identity

import (
	"context"
	"time"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"github.com/dizzrt/dauth/internal/domain/identity/repo"
	"github.com/dizzrt/dauth/internal/infra/foundation"
	identity_model "github.com/dizzrt/dauth/internal/infra/repo/model/identity"
)

var _ repo.UserRepo = (*UserRepoImpl)(nil)

type UserRepoImpl struct {
	*foundation.BaseDB
}

func NewUserRepoImpl(base *foundation.BaseDB) repo.UserRepo {
	return &UserRepoImpl{
		BaseDB: base,
	}
}

func (impl *UserRepoImpl) CreateUser(ctx context.Context, user *entity.User) (uint32, error) {
	model := &identity_model.User{
		Email:         user.Email,
		Username:      user.Username,
		Password:      user.Password,
		Status:        uint(identity.UserStatus_ENABLED),
		LastLoginTime: time.Now(),
	}

	db := impl.WithContext(ctx)
	if err := db.Create(&model).Error; err != nil {
		return 0, impl.WrapError(err)
	}

	return uint32(model.ID), nil
}

func (impl *UserRepoImpl) GetUserByID(ctx context.Context, uid uint32) (*entity.User, error) {
	var model *identity_model.User
	db := impl.WithContext(ctx)
	if err := db.Where("id = ?", uid).First(&model).Error; err != nil {
		return nil, impl.WrapError(err)
	}

	return model.ToEntity(), nil
}
