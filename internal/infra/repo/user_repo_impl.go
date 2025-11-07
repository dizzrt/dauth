package repo

import (
	"context"
	"time"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"github.com/dizzrt/dauth/internal/domain/identity/repo"
	"github.com/dizzrt/dauth/internal/infra/repo/model"
	"gorm.io/gorm"
)

var _ repo.UserRepo = (*UserRepoImpl)(nil)

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepoImpl(db *gorm.DB) repo.UserRepo {
	return &UserRepoImpl{db: db}
}

func (impl *UserRepoImpl) CreateUser(ctx context.Context, user *entity.User) (uint32, error) {
	model := &model.User{
		Email:         user.Email,
		Password:      user.Password,
		Username:      user.Username,
		Status:        0,
		LastLoginTime: time.Now(),
	}

	db := impl.db.WithContext(ctx)
	if err := db.Create(model).Error; err != nil {
		return 0, err
	}

	return uint32(model.ID), nil
}

func (impl *UserRepoImpl) GetUserByID(ctx context.Context, uid uint32) (*entity.User, error) {
	var model model.User
	db := impl.db.WithContext(ctx)
	if err := db.Where("id = ?", uid).First(&model).Error; err != nil {
		return nil, err
	}

	return &entity.User{
		ID:            model.ID,
		Email:         model.Email,
		Username:      model.Username,
		Password:      model.Password,
		Status:        model.Status,
		LastLoginTime: model.LastLoginTime,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		DeletedAt:     model.DeletedAt,
	}, nil
}

func (impl *UserRepoImpl) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var model model.User
	db := impl.db.WithContext(ctx)
	if err := db.Where("email = ?", email).First(&model).Error; err != nil {
		return nil, err
	}

	return &entity.User{
		ID:            model.ID,
		Email:         model.Email,
		Username:      model.Username,
		Password:      model.Password,
		Status:        model.Status,
		LastLoginTime: model.LastLoginTime,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		DeletedAt:     model.DeletedAt,
	}, nil
}

func (impl *UserRepoImpl) UpdateUserPassword(ctx context.Context, uid uint32, password string) error {
	db := impl.db.WithContext(ctx)
	if err := db.Model(&model.User{}).Where("id = ?", uid).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}

func (impl *UserRepoImpl) UpdateUserStatus(ctx context.Context, uid uint32, status identity.UserStatus) error {
	db := impl.db.WithContext(ctx)
	if err := db.Model(&model.User{}).Where("id = ?", uid).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
