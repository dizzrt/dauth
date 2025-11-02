package handler

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/user"
	"github.com/dizzrt/dauth/internal/application"
	"github.com/dizzrt/ellie/log"
)

type UserHandler struct {
	user.UnimplementedUserServiceServer

	userApp application.UserApplication
}

func NewUserHandler(userApp application.UserApplication) *UserHandler {
	return &UserHandler{userApp: userApp}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	log.CtxInfof(ctx, "[CreateUser] handler req: %v", req)
	return h.userApp.CreateUser(ctx, req)
}

func (h *UserHandler) GetUserByID(ctx context.Context, req *user.GetUserByIDRequest) (*user.GetUserByIDResponse, error) {
	log.CtxInfof(ctx, "[GetUserByID] handler req: %v", req)
	return h.userApp.GetUserByID(ctx, req)
}
