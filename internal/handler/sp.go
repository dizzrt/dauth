package handler

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/sp"
	"github.com/dizzrt/dauth/internal/application"
)

var _ sp.ServiceProviderServiceServer = (*ServiceProviderHandler)(nil)

type ServiceProviderHandler struct {
	sp.UnimplementedServiceProviderServiceServer

	spApp application.ServiceProviderApplication
}

func NewServiceProviderHandler(spApp application.ServiceProviderApplication) *ServiceProviderHandler {
	return &ServiceProviderHandler{
		spApp: spApp,
	}
}

func (handler *ServiceProviderHandler) CreateServiceProvider(ctx context.Context, req *sp.CreateServiceProviderRequest) (*sp.CreateServiceProviderResponse, error) {
	return handler.spApp.CreateServiceProvider(ctx, req)
}

func (handler *ServiceProviderHandler) GetServiceProvider(ctx context.Context, req *sp.GetServiceProviderRequest) (*sp.GetServiceProviderResponse, error) {
	return handler.spApp.GetServiceProvider(ctx, req)
}

func (handler *ServiceProviderHandler) ValidateServiceProvider(ctx context.Context, req *sp.ValidateServiceProviderRequest) (*sp.ValidateServiceProviderResponse, error) {
	return handler.spApp.ValidateServiceProvider(ctx, req)
}
