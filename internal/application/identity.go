package application

import (
	"context"
	"fmt"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/domain/identity/biz"
	"github.com/dizzrt/dauth/internal/domain/identity/dto"
	"github.com/dizzrt/ellie/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ IdentityApplication = (*identityApplication)(nil)

type IdentityApplication interface {
	CreateUser(context.Context, *identity.CreateUserRequest) (*identity.CreateUserResponse, error)
	GetUser(context.Context, *identity.GetUserRequest) (*identity.GetUserResponse, error)
}

type identityApplication struct {
	userBiz biz.UserBiz
}

func NewIdentityApplication(userBiz biz.UserBiz) IdentityApplication {
	return &identityApplication{
		userBiz: userBiz,
	}
}

func (app *identityApplication) CreateUser(ctx context.Context, req *identity.CreateUserRequest) (*identity.CreateUserResponse, error) {
	pwd := req.GetPassword()
	if pwd == "" {
		return nil, fmt.Errorf("password can not be empty")
	}

	username := req.GetUsername()
	if username == "" {
		return nil, fmt.Errorf("username can not be empty")
	}

	dto := &dto.CreateUserDTO{
		Password: req.GetPassword(),
		Username: req.GetUsername(),
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
		Avatar:   req.Email,
	}

	user, err := app.userBiz.CreateUser(ctx, dto)
	if err != nil {
		log.CtxErrorf(ctx, "")
		return nil, err
	}

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

	resp := &identity.CreateUserResponse{
		User: u,
		// BaseResp: ,
	}

	return resp, nil
}

func (app *identityApplication) GetUser(ctx context.Context, req *identity.GetUserRequest) (*identity.GetUserResponse, error) {
	uid := req.GetUid()
	if uid == 0 {
		return nil, fmt.Errorf("invalid uid")
	}

	user, err := app.userBiz.GetUser(ctx, uid)
	if err != nil {
		log.CtxErrorf(ctx, "get user by id failed, err: %v", err)
		return nil, err
	}

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

	resp := &identity.GetUserResponse{
		User: u,
		// BaseResp: ,
	}

	return resp, nil
}
