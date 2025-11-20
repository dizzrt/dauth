package repo

import (
	"context"
	"time"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"github.com/dizzrt/dauth/internal/domain/identity/repo"
	"github.com/dizzrt/dauth/internal/infra/foundation"
	"github.com/dizzrt/dauth/internal/infra/repo/model"
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
	model := &model.User{
		Email:         user.Email,
		Username:      user.Username,
		Password:      user.Password,
		Status:        uint(identity.UserStatus_ACTIVE),
		LastLoginTime: time.Now(),
	}

	db := impl.WithContext(ctx)
	if err := db.Create(&model).Error; err != nil {
		return 0, err
	}

	return uint32(model.ID), nil
}

func (impl *UserRepoImpl) GetUserByID(ctx context.Context, uid uint32) (*entity.User, error) {
	var model *model.User
	db := impl.WithContext(ctx)
	if err := db.Where("id = ?", uid).First(&model).Error; err != nil {
		return nil, err
	}

	return model.ToEntity()
}

func (impl *UserRepoImpl) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var model *model.User

	db := impl.WithContext(ctx)
	if err := db.Where("email = ?", email).
		First(&model).Error; err != nil {
		return nil, err
	}

	return model.ToEntity()
}

func (impl *UserRepoImpl) UpdateUserPassword(ctx context.Context, uid uint32, password string) error {
	db := impl.WithContext(ctx)
	err := db.Model(&model.User{}).Where("id = ?", uid).
		Update("password", password).Error

	return err
}

func (impl *UserRepoImpl) UpdateUserStatus(ctx context.Context, uid uint32, status identity.UserStatus) error {
	mStatus := uint(status)
	db := impl.WithContext(ctx)
	err := db.Model(&model.User{}).
		Where("id = ?", uid).
		Update("status", mStatus).Error

	return err
}
