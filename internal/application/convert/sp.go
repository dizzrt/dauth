package convert

import (
	"github.com/dizzrt/dauth/api/gen/sp"
	"github.com/dizzrt/dauth/internal/domain/sp/entity"
)

func SPEntityFromCreateSPRequest(req *sp.CreateServiceProviderRequest) *entity.ServiceProvider {
	return &entity.ServiceProvider{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Secret:      req.GetSecret(),
		RedirectURI: req.GetRedirectUri(),
	}
}
