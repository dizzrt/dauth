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

func (handler *ClientHandler) CreateClient(ctx context.Context, req *client.CreateClientRequest) (*client.CreateClientResponse, error) {
	return handler.clientApp.CreateClient(ctx, req)
}

func (handler *ClientHandler) GetClient(ctx context.Context, req *client.GetClientRequest) (*client.GetClientResponse, error) {
	return handler.clientApp.GetClient(ctx, req)
}

func (handler *ClientHandler) ValidateClient(ctx context.Context, req *client.ValidateClientRequest) (*client.ValidateClientResponse, error) {
	return handler.clientApp.ValidateClient(ctx, req)
}
