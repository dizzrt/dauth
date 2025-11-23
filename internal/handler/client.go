package handler

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/client"
	"github.com/dizzrt/dauth/internal/application"
)

var _ client.ClientServiceServer = (*ClientHandler)(nil)

type ClientHandler struct {
	client.UnimplementedClientServiceServer

	clientApp application.ClientApplication
}

func NewClientHandler(clientApp application.ClientApplication) *ClientHandler {
	return &ClientHandler{
		clientApp: clientApp,
	}
}

func (handler *ClientHandler) Create(ctx context.Context, req *client.CreateRequest) (*client.CreateResponse, error) {
	return handler.clientApp.Create(ctx, req)
}

func (handler *ClientHandler) Validate(ctx context.Context, req *client.ValidateRequest) (*client.ValidateResponse, error) {
	return handler.clientApp.Validate(ctx, req)
}
