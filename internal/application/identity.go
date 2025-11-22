package application

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/application/convert"
	"github.com/dizzrt/dauth/internal/domain/identity/biz"
	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"github.com/dizzrt/dauth/internal/infra/rpc"
	"github.com/dizzrt/dauth/internal/infra/utils"
	"github.com/dizzrt/ellie/log"
)

var _ IdentityApplication = (*identityApplication)(nil)

type IdentityApplication interface {
	Login(context.Context, *identity.LoginRequest) (*identity.LoginResponse, error)
	Authenticate(context.Context, *identity.AuthenticateRequest) (*identity.AuthenticateResponse, error)
	CreateUser(context.Context, *identity.CreateUserRequest) (*identity.CreateUserResponse, error)
	GetUser(context.Context, *identity.GetUserRequest) (*identity.GetUserResponse, error)
	UpdateUserStatus(context.Context, *identity.UpdateUserStatusRequest) (*identity.UpdateUserStatusResponse, error)
	UpdateUserPassword(context.Context, *identity.UpdateUserPasswordRequest) (*identity.UpdateUserPasswordResponse, error)
}

type identityApplication struct {
	userBiz biz.UserBiz
	roleBiz biz.RoleBiz
}

func NewIdentityApplication(userBiz biz.UserBiz, roleBiz biz.RoleBiz) IdentityApplication {
	return &identityApplication{
		userBiz: userBiz,
		roleBiz: roleBiz,
	}
}

func (app *identityApplication) Login(ctx context.Context, req *identity.LoginRequest) (*identity.LoginResponse, error) {
	var account string // only support email yet
	if account = req.GetAccount(); account == "" {
		return nil, identity.ErrorInvalidEmail("email can not be empty")
	}

	if err := utils.Validate().Var(account, "email"); err != nil {
		return nil, identity.ErrorInvalidEmail("malformed email").WithCause(err)
	}

	var pwd string
	if pwd = req.GetPassword(); pwd == "" {
		return nil, identity.ErrorInvalidPassword("password can not be empty")
	}

	// authenticate user
	user, err := app.userBiz.Authenticate(ctx, account, pwd)
	if err != nil {
		return nil, err
	}

	// update last login time
	if err := app.userBiz.UpdateLastLoginTime(ctx, user.ID); err != nil {
		log.CtxErrorf(ctx, "user `%d` has logged in successfully, but failed to update the last login time; err: %v", user.ID, err)
	}

	return &identity.LoginResponse{
		User:     convert.UserToIdentityUser(user),
		BaseResp: rpc.Success(),
	}, nil
}

func (app *identityApplication) Authenticate(ctx context.Context, req *identity.AuthenticateRequest) (*identity.AuthenticateResponse, error) {
	var account string // only support email yet
	if account = req.GetAccount(); account == "" {
		return nil, identity.ErrorInvalidEmail("email can not be empty")
	}

	if err := utils.Validate().Var(account, "email"); err != nil {
		return nil, identity.ErrorInvalidEmail("malformed email").WithCause(err)
	}

	var pwd string
	if pwd = req.GetPassword(); pwd == "" {
		return nil, identity.ErrorInvalidPassword("password can not be empty")
	}

	user, err := app.userBiz.Authenticate(ctx, account, pwd)
	if err != nil {
		return nil, err
	}

	return &identity.AuthenticateResponse{
		User:     convert.UserToIdentityUser(user),
		BaseResp: rpc.Success(),
	}, nil
}

func (app *identityApplication) CreateUser(ctx context.Context, req *identity.CreateUserRequest) (*identity.CreateUserResponse, error) {
	user := &entity.User{
		Email:    req.GetEmail(),
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Status:   identity.User_ACTIVE,
	}

	// validate user
	if err := utils.Validate().Struct(user); err != nil {
		return nil, identity.ErrorInvalidParams("").WithCause(err)
	}

	uid, err := app.userBiz.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &identity.CreateUserResponse{
		Id:       uid,
		BaseResp: rpc.Success(),
	}, nil
}

func (app *identityApplication) GetUser(ctx context.Context, req *identity.GetUserRequest) (*identity.GetUserResponse, error) {
	uid := req.GetId()
	user, err := app.userBiz.GetUserByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return &identity.GetUserResponse{
		User:     convert.UserToIdentityUser(user),
		BaseResp: rpc.Success(),
	}, nil
}

func (app *identityApplication) UpdateUserStatus(ctx context.Context, req *identity.UpdateUserStatusRequest) (*identity.UpdateUserStatusResponse, error) {
	uid := req.GetId()
	status := req.GetStatus()

	err := app.userBiz.UpdateUserStatus(ctx, uid, status)
	if err != nil {
		return nil, err
	}

	return &identity.UpdateUserStatusResponse{
		BaseResp: rpc.Success(),
	}, nil
}

func (app *identityApplication) UpdateUserPassword(ctx context.Context, req *identity.UpdateUserPasswordRequest) (*identity.UpdateUserPasswordResponse, error) {
	uid := req.GetId()
	pwd := req.GetPassword()
	if pwd == "" {
		return nil, identity.ErrorInvalidPassword("password can not be empty")
	}

	err := app.userBiz.UpdateUserPassword(ctx, uid, pwd)
	if err != nil {
		return nil, err
	}

	return &identity.UpdateUserPasswordResponse{
		BaseResp: rpc.Success(),
	}, nil
}
