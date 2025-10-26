package handler

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/example"
	"github.com/dizzrt/dauth/internal/application"
)

type ExampleHandler struct {
	example.UnimplementedExampleServer

	app application.ExampleApplication
}

func NewExampleHandler(app application.ExampleApplication) *ExampleHandler {
	return &ExampleHandler{app: app}
}

func (h *ExampleHandler) Hello(ctx context.Context, req *example.HelloRequest) (*example.HelloResponse, error) {
	msg, err := h.app.Hello(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &example.HelloResponse{
		Message: msg,
	}, nil
}
