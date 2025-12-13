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
		Status:        uint(identity.User_ACTIVE),
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

func (impl *UserRepoImpl) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var model *identity_model.User

	db := impl.WithContext(ctx)
	if err := db.Where("email = ?", email).
		First(&model).Error; err != nil {
		return nil, impl.WrapError(err)
	}

	return model.ToEntity(), nil
}

func (impl *UserRepoImpl) UpdateUserPassword(ctx context.Context, uid uint32, password string) error {
	db := impl.WithContext(ctx)
	err := db.Model(&identity_model.User{}).Where("id = ?", uid).
		Update("password", password).Error

	return impl.WrapError(err)
}

func (impl *UserRepoImpl) UpdateUserStatus(ctx context.Context, uid uint32, status identity.User_Status) error {
	mStatus := uint(status)
	db := impl.WithContext(ctx)
	err := db.Model(&identity_model.User{}).
		Where("id = ?", uid).
		Update("status", mStatus).Error

	return impl.WrapError(err)
}

func (impl *UserRepoImpl) UpdateLastLoginTime(ctx context.Context, uid uint32, lastLoginTime time.Time) error {
	db := impl.WithContext(ctx)
	err := db.Model(&identity_model.User{}).
		Where("id = ?", uid).
		Update("last_login_time", lastLoginTime).Error

	return impl.WrapError(err)
}
