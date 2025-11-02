package handler

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/user"
	"github.com/dizzrt/dauth/internal/application"
)

type UserHandler struct {
	user.UnimplementedUserServiceServer

	userApp application.UserApplication
}

func NewUserHandler(userApp application.UserApplication) *UserHandler {
	return &UserHandler{userApp: userApp}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	return h.userApp.CreateUser(ctx, req)
}

func (h *UserHandler) GetUserByID(ctx context.Context, req *user.GetUserByIDRequest) (*user.GetUserResponse, error) {
	return h.userApp.GetUserByID(ctx, req)
}

func (h *UserHandler) GetUserByEmail(ctx context.Context, req *user.GetUserByEmailRequest) (*user.GetUserResponse, error) {
	return h.userApp.GetUserByEmail(ctx, req)
}

func (h *UserHandler) UserLogin(ctx context.Context, req *user.UserLoginRequest) (*user.GetUserResponse, error) {
	return h.userApp.UserLogin(ctx, req)
}
