package biz

import (
	"context"
	"errors"

	api_user "github.com/dizzrt/dauth/api/gen/user"
	"github.com/dizzrt/dauth/internal/domain/user/entity"
	"github.com/dizzrt/dauth/internal/domain/user/repo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var _ UserBiz = (*userBiz)(nil)

type UserBiz interface {
	// create user, return uid when success
	CreateUser(ctx context.Context, user *entity.User) (uint, error)
	GetUserByID(ctx context.Context, uid uint32) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UserLogin(ctx context.Context, account string, password string) (*entity.User, error)
}

type userBiz struct {
	userRepo repo.UserRepo
}

func NewUserBiz(userRepo repo.UserRepo) UserBiz {
	return &userBiz{
		userRepo: userRepo,
	}
}

func (biz *userBiz) CreateUser(ctx context.Context, user *entity.User) (uint, error) {
	pwd := user.Password
	pwdBytes := []byte(pwd)
	hashBytes, err := bcrypt.GenerateFromPassword(pwdBytes, 10)
	if err != nil {
		return 0, err
	}

	// save hash password
	user.Password = string(hashBytes)
	uid, err := biz.userRepo.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}

	return uid, nil
}

func (biz *userBiz) GetUserByID(ctx context.Context, uid uint32) (*entity.User, error) {
	user, err := biz.userRepo.GetUserByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (biz *userBiz) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := biz.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// only login by email yet
func (biz *userBiz) UserLogin(ctx context.Context, account string, password string) (*entity.User, error) {
	user, err := biz.GetUserByEmail(ctx, account)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, api_user.ErrorUserNotFound("user not found")
		}

		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, api_user.ErrorInvalidPassword("wrong password")
	}

	return user, nil
}
