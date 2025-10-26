package application

import (
	"context"

	"github.com/dizzrt/dauth/internal/domain/biz"
)

type ExampleApplication interface {
	Hello(ctx context.Context, name string) (string, error)
}

type exampleApplication struct {
	biz biz.ExampleBiz
}

func NewExampleApplication(biz biz.ExampleBiz) ExampleApplication {
	return &exampleApplication{biz: biz}
}

func (app *exampleApplication) Hello(ctx context.Context, name string) (string, error) {
	return app.biz.Hello(ctx, name)
}
