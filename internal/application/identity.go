package application

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/common"
	"github.com/dizzrt/dauth/api/gen/errdef"
	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/application/convert"
	"github.com/dizzrt/dauth/internal/domain/identity/biz"
	"github.com/dizzrt/dauth/internal/domain/identity/dto"
	"github.com/dizzrt/dauth/internal/infra/rpc"
	"github.com/dizzrt/ellie/log"
)

var _ IdentityApplication = (*identityApplication)(nil)

type IdentityApplication interface {
	CreateUser(context.Context, *identity.CreateUserRequest) (*identity.CreateUserResponse, error)
	GetUser(context.Context, *identity.GetUserRequest) (*identity.GetUserResponse, error)
	GetUserByName(context.Context, *identity.GetUserByNameRequest) (*identity.GetUserByNameResponse, error)
	ListUsers(context.Context, *identity.ListUsersRequest) (*identity.ListUsersResponse, error)
	UpdateUserStatus(context.Context, *identity.UpdateUserStatusRequest) (*identity.UpdateUserStatusResponse, error)
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
	username := req.GetUsername()
	if username == "" {
		return nil, errdef.InvalidArgument().WithMessage("username can not be empty")
	}

	pwd := req.GetPassword()
	if pwd == "" {
		return nil, errdef.InvalidArgument().WithMessage("password can not be empty")
	}

	dto := &dto.CreateUserDTO{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
		Avatar:   req.Email,
	}

	user, err := app.userBiz.CreateUser(ctx, dto)
	if err != nil {
		if errdef.IsDuplicatedKey(err) {
			return nil, errdef.UserExist().WithMessage("user already exist")
		}

		log.CtxErrorf(ctx, "create user failed, err: %v", err)
		return nil, err
	}

	resp := &identity.CreateUserResponse{
		User:     convert.ToIdentityUser(user),
		BaseResp: rpc.Success(),
	}

	return resp, nil
}

func (app *identityApplication) GetUser(ctx context.Context, req *identity.GetUserRequest) (*identity.GetUserResponse, error) {
	uid := req.GetUid()
	if uid == 0 {
		return nil, errdef.InvalidArgument().WithMessage("invalid uid")
	}

	user, err := app.userBiz.GetUser(ctx, uid)
	if err != nil {
		log.CtxErrorf(ctx, "get user by id failed, err: %v", err)
		return nil, err
	}

	resp := &identity.GetUserResponse{
		User:     convert.ToIdentityUser(user),
		BaseResp: rpc.Success(),
	}

	return resp, nil
}

func (app *identityApplication) GetUserByName(ctx context.Context, req *identity.GetUserByNameRequest) (*identity.GetUserByNameResponse, error) {
	username := req.GetUsername()
	if username == "" {
		return nil, errdef.InvalidArgument().WithMessage("username can not be empty")
	}

	user, err := app.userBiz.GetUserByName(ctx, username)
	if err != nil {
		log.CtxErrorf(ctx, "get user by name failed, err: %v", err)
		return nil, err
	}

	resp := &identity.GetUserByNameResponse{
		User:     convert.ToIdentityUser(user),
		BaseResp: rpc.Success(),
	}

	return resp, nil
}

func (app *identityApplication) ListUsers(ctx context.Context, req *identity.ListUsersRequest) (*identity.ListUsersResponse, error) {
	page := req.GetPagination().GetPage()
	if page <= 0 {
		page = 1
	}

	size := req.GetPagination().GetSize()
	if size <= 0 {
		size = 20
	}

	users, total, err := app.userBiz.ListUsers(ctx, page, size)
	if err != nil {
		log.CtxErrorf(ctx, "list users failed, err: %v", err)
		return nil, err
	}

	resp := &identity.ListUsersResponse{
		Users: convert.ToIdentityUsers(users),
		Pagination: &common.Pagination{
			Page:  page,
			Size:  size,
			Total: total,
		},
	}

	return resp, nil
}

func (app *identityApplication) UpdateUserStatus(ctx context.Context, req *identity.UpdateUserStatusRequest) (*identity.UpdateUserStatusResponse, error) {
	uid := req.GetUid()
	if uid == 0 {
		return nil, errdef.InvalidArgument().WithMessage("invalid uid")
	}

	status := req.GetStatus()
	if status == identity.UserStatus_USER_STATUS_UNSPECIFIED {
		return nil, errdef.InvalidArgument().WithMessage("unknown user status '%v'", status)
	}

	newStatus, err := app.userBiz.UpdateUserStatus(ctx, uid, status)
	if err != nil {
		log.CtxErrorf(ctx, "update user status failed, err: %v", err)
		return nil, err
	}

	resp := &identity.UpdateUserStatusResponse{
		Status: newStatus,
	}

	return resp, nil
}
