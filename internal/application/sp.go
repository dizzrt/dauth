package application

import (
	"context"
	"errors"

	"github.com/dizzrt/dauth/api/gen/errdef"
	"github.com/dizzrt/dauth/api/gen/sp"
	"github.com/dizzrt/dauth/internal/application/convert"
	"github.com/dizzrt/dauth/internal/domain/sp/biz"
	"github.com/dizzrt/dauth/internal/infra/rpc"
)

var _ ServiceProviderApplication = (*serviceProviderApplication)(nil)

type ServiceProviderApplication interface {
	CreateServiceProvider(ctx context.Context, req *sp.CreateServiceProviderRequest) (*sp.CreateServiceProviderResponse, error)
	GetServiceProvider(ctx context.Context, req *sp.GetServiceProviderRequest) (*sp.GetServiceProviderResponse, error)
	ListServiceProvider(ctx context.Context, req *sp.ListServiceProviderRequest) (*sp.ListServiceProviderResponse, error)
	ValidateServiceProvider(ctx context.Context, req *sp.ValidateServiceProviderRequest) (*sp.ValidateServiceProviderResponse, error)
}

type serviceProviderApplication struct {
	spBiz biz.ServiceProviderBiz
}

func NewServiceProviderApplication(spBiz biz.ServiceProviderBiz) ServiceProviderApplication {
	return &serviceProviderApplication{
		spBiz: spBiz,
	}
}

func (app *serviceProviderApplication) CreateServiceProvider(ctx context.Context, req *sp.CreateServiceProviderRequest) (*sp.CreateServiceProviderResponse, error) {
	spEntity := convert.SPEntityFromCreateSPRequest(req)
	if spEntity.Secret == "" {
		return nil, errors.New("secret is required")
	}

	scopeIDs := req.GetScopes()
	spID, err := app.spBiz.CreateServiceProvider(ctx, spEntity, scopeIDs)
	if err != nil {
		return nil, err
	}

	return &sp.CreateServiceProviderResponse{
		SpId:     spID,
		BaseResp: rpc.Success(),
	}, nil
}

func (app *serviceProviderApplication) GetServiceProvider(ctx context.Context, req *sp.GetServiceProviderRequest) (*sp.GetServiceProviderResponse, error) {
	spID := req.GetSpId()
	if spID == 0 {
		return nil, nil
	}

	spEntity, err := app.spBiz.GetServiceProvider(ctx, spID)
	if err != nil {
		return nil, err
	}

	if spEntity.DeletedAt.Valid {
		spEntity.Status = sp.ServiceProvider_DELETED
	}

	return &sp.GetServiceProviderResponse{
		Sp: &sp.ServiceProvider{
			Id:          spEntity.ID,
			Name:        spEntity.Name,
			Description: spEntity.Description,
			RedirectUri: spEntity.RedirectURI,
			Status:      spEntity.Status,
			CreatedAt:   spEntity.CreatedAt.Unix(),
			UpdatedAt:   spEntity.UpdatedAt.Unix(),
		},
		BaseResp: rpc.Success(),
	}, nil
}

func (app *serviceProviderApplication) ListServiceProvider(ctx context.Context, req *sp.ListServiceProviderRequest) (*sp.ListServiceProviderResponse, error) {
	dtoReq := convert.ListServiceProviderReqFromPB(req)
	if dtoReq.Page < 1 {
		return nil, errdef.InvalidParamsWithMsg("page must be greater than 0")
	}

	if dtoReq.PageSize < 1 {
		return nil, errdef.InvalidParamsWithMsg("page_size must be greater than 0")
	}

	spList, total, err := app.spBiz.ListServiceProvider(ctx, dtoReq)
	if err != nil {
		return nil, err
	}

	respList := make([]*sp.ServiceProvider, 0, len(spList))
	for _, p := range spList {
		temp := &sp.ServiceProvider{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			RedirectUri: p.RedirectURI,
			Status:      p.Status,
			CreatedAt:   p.CreatedAt.Unix(),
			UpdatedAt:   p.UpdatedAt.Unix(),
		}

		if p.DeletedAt.Valid {
			temp.Status = sp.ServiceProvider_DELETED
		}
		respList = append(respList, temp)
	}

	return &sp.ListServiceProviderResponse{
		SpList:   respList,
		Total:    total,
		BaseResp: rpc.Success(),
	}, nil
}

func (app *serviceProviderApplication) ValidateServiceProvider(ctx context.Context, req *sp.ValidateServiceProviderRequest) (*sp.ValidateServiceProviderResponse, error) {
	spID := req.GetSpId()
	if spID == 0 {
		return nil, errors.New("sp_id is required")
	}

	isOK, reason, err := app.spBiz.ValidateServiceProvider(ctx, spID, req.GetScope())
	if err != nil {
		return nil, err
	}

	return &sp.ValidateServiceProviderResponse{
		IsOk:     isOK,
		Reason:   reason,
		BaseResp: rpc.Success(),
	}, nil
}
