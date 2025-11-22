package handler

import (
	"github.com/dizzrt/dauth/api/gen/client"
	"github.com/dizzrt/dauth/internal/application"
)

type ClientHandler struct {
	client.UnimplementedClientServiceServer

	clientApp application.ClientApplication
}

func NewClientHandler(clientApp application.ClientApplication) *ClientHandler {
	return &ClientHandler{
		clientApp: clientApp,
	}
}
