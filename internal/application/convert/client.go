package convert

import (
	"github.com/dizzrt/dauth/api/gen/client"
	"github.com/dizzrt/dauth/internal/domain/client/entity"
)

func ClientEntityFromCreateClientRequest(req *client.CreateClientRequest) *entity.Client {
	return &entity.Client{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Secret:      req.GetSecret(),
		RedirectURI: req.GetRedirectUri(),
	}
}
