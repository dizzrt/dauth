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
	GetUserByID(context.Context, *user.GetUserByIDRequest) (*user.GetUserByIDResponse, error)
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
	log.CtxInfof(ctx, "[CreateUser] app req: %v", req)

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

func (app *userApplication) GetUserByID(ctx context.Context, req *user.GetUserByIDRequest) (*user.GetUserByIDResponse, error) {
	log.CtxInfof(ctx, "[GetUserByID] app req: %v", req)

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

	return &user.GetUserByIDResponse{
		User: resp,
	}, err
}
