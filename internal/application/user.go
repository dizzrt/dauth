package application

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/user"
	"github.com/dizzrt/dauth/internal/domain/user/biz"
	"github.com/dizzrt/dauth/internal/domain/user/entity"
	"github.com/dizzrt/ellie/log"
)

var _ UserApplication = (*userApplication)(nil)

type UserApplication interface {
	// create user, return uid when success
	CreateUser(context.Context, *user.CreateUserRequest) (*user.CreateUserResponse, error)
	GetUserByID(context.Context, *user.GetUserByIDRequest) (*user.GetUserResponse, error)
	GetUserByEmail(context.Context, *user.GetUserByEmailRequest) (*user.GetUserResponse, error)
	UserLogin(context.Context, *user.UserLoginRequest) (*user.GetUserResponse, error)
}

type userApplication struct {
	userBiz biz.UserBiz
}

func NewUserApplication(userBiz biz.UserBiz) UserApplication {
	return &userApplication{
		userBiz: userBiz,
	}
}

func (app *userApplication) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	userEntity := &entity.User{
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
	}

	uid, err := app.userBiz.CreateUser(ctx, userEntity)
	if err != nil {
		log.CtxErrorf(ctx, "[CreateUser] failed, err: %v", err)
	}

	return &user.CreateUserResponse{
		Id: uint32(uid),
	}, err
}

func (app *userApplication) GetUserByID(ctx context.Context, req *user.GetUserByIDRequest) (*user.GetUserResponse, error) {
	userEntity, err := app.userBiz.GetUserByID(ctx, req.GetId())
	if err != nil {
		log.CtxErrorf(ctx, "[GetUserByID] failed, err: %v", err)
		return nil, err
	}

	resp := &user.User{
		Id:          uint32(userEntity.ID),
		Email:       userEntity.Email,
		Username:    userEntity.Username,
		Status:      uint32(userEntity.Status),
		LastLoginAt: uint64(userEntity.LastLoginTime.UnixMilli()),
		CreatedAt:   uint64(userEntity.CreatedAt.UnixMilli()),
		UpdatedAt:   uint64(userEntity.UpdatedAt.UnixMilli()),
	}

	if userEntity.DeletedAt.Valid {
		resp.DeletedAt = uint64(userEntity.DeletedAt.Time.UnixMilli())
	}

	return &user.GetUserResponse{
		User: resp,
	}, err
}

func (app *userApplication) GetUserByEmail(ctx context.Context, req *user.GetUserByEmailRequest) (*user.GetUserResponse, error) {
	userEntity, err := app.userBiz.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		log.CtxErrorf(ctx, "[GetUserByEmail] failed, err: %v", err)
		return nil, err
	}

	resp := &user.User{
		Id:          uint32(userEntity.ID),
		Email:       userEntity.Email,
		Username:    userEntity.Username,
		Status:      uint32(userEntity.Status),
		LastLoginAt: uint64(userEntity.LastLoginTime.UnixMilli()),
		CreatedAt:   uint64(userEntity.CreatedAt.UnixMilli()),
		UpdatedAt:   uint64(userEntity.UpdatedAt.UnixMilli()),
	}

	if userEntity.DeletedAt.Valid {
		resp.DeletedAt = uint64(userEntity.DeletedAt.Time.UnixMilli())
	}

	return &user.GetUserResponse{
		User: resp,
	}, err
}

func (app *userApplication) UserLogin(ctx context.Context, req *user.UserLoginRequest) (*user.GetUserResponse, error) {
	userEntity, err := app.userBiz.UserLogin(ctx, req.GetAccount(), req.GetPassword())
	if err != nil {
		if !user.IsInvalidPassword(err) && !user.IsUserNotFound(err) {
			log.CtxErrorf(ctx, "[UserLogin] failed, err: %v", err)
		}

		return nil, err
	}

	return &user.GetUserResponse{
		User: &user.User{
			Id:          uint32(userEntity.ID),
			Email:       userEntity.Email,
			Username:    userEntity.Username,
			Status:      uint32(userEntity.Status),
			LastLoginAt: uint64(userEntity.LastLoginTime.UnixMilli()),
			CreatedAt:   uint64(userEntity.CreatedAt.UnixMilli()),
			UpdatedAt:   uint64(userEntity.UpdatedAt.UnixMilli()),
		},
	}, err
}
