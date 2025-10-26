package biz

import (
	"context"

	"github.com/dizzrt/ellie/log"
)

type ExampleBiz interface {
	Hello(ctx context.Context, name string) (string, error)
}

type exampleBiz struct {
}

func NewExampleBiz() ExampleBiz {
	return &exampleBiz{}
}

func (biz *exampleBiz) Hello(ctx context.Context, name string) (string, error) {
	log.Error("test error")
	return "hello " + name + "!!!", nil
}
