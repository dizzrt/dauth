package repo

import (
	"context"
	"time"

	"github.com/dizzrt/dauth/internal/domain/user/entity"
	user_repo "github.com/dizzrt/dauth/internal/domain/user/repo"
	"github.com/dizzrt/dauth/internal/infra/common"
	"github.com/dizzrt/dauth/internal/infra/repo/model"
)

var _ user_repo.UserRepo = (*UserRepoImpl)(nil)

type UserRepoImpl struct {
}

func NewUserRepoImpl() user_repo.UserRepo {
	return &UserRepoImpl{}
}

func (impl *UserRepoImpl) CreateUser(ctx context.Context, user *entity.User) (uint, error) {
	model := &model.User{
		Email:         user.Email,
		Password:      user.Password,
		Username:      user.Username,
		Status:        0,
		LastLoginTime: time.Now(),
	}

	db := common.DB().WithContext(ctx)
	if err := db.Create(model).Error; err != nil {
		return 0, err
	}

	return model.ID, nil
}

func (impl *UserRepoImpl) GetUserByID(ctx context.Context, uid uint32) (*entity.User, error) {
	var model model.User
	db := common.DB().WithContext(ctx)
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
	db := common.DB().WithContext(ctx)
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
