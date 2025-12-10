package convert

import (
	"github.com/dizzrt/dauth/api/gen/sp"
	"github.com/dizzrt/dauth/internal/domain/sp/dto"
	"github.com/dizzrt/dauth/internal/domain/sp/entity"
)

func ListServiceProviderReqFromPB(req *sp.ListServiceProviderRequest) *dto.ListServiceProviderRequest {
	return &dto.ListServiceProviderRequest{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
	}
}

func SPEntityFromCreateSPRequest(req *sp.CreateServiceProviderRequest) *entity.ServiceProvider {
	return &entity.ServiceProvider{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Secret:      req.GetSecret(),
		RedirectURI: req.GetRedirectUri(),
	}
}
