package iface

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/example"
)

type ExampleHandler struct {
	example.UnimplementedExampleServer
}

func NewExampleHandler() *ExampleHandler {
	return &ExampleHandler{}
}

func (h *ExampleHandler) Hello(ctx context.Context, req *example.HelloRequest) (*example.HelloResponse, error) {
	return &example.HelloResponse{
		Message: "hello " + req.Name,
	}, nil
}
