package application

import (
	"context"
	"errors"

	"github.com/dizzrt/dauth/api/gen/client"
	"github.com/dizzrt/dauth/internal/application/convert"
	"github.com/dizzrt/dauth/internal/domain/client/biz"
	"github.com/dizzrt/dauth/internal/infra/rpc"
)

var _ ClientApplication = (*clientApplication)(nil)

type ClientApplication interface {
	CreateClient(ctx context.Context, req *client.CreateClientRequest) (*client.CreateClientResponse, error)
	ValidateClient(ctx context.Context, req *client.ValidateClientRequest) (*client.ValidateClientResponse, error)
}

type clientApplication struct {
	clientBiz biz.ClientBiz
}

func NewClientApplication(clientBiz biz.ClientBiz) ClientApplication {
	return &clientApplication{
		clientBiz: clientBiz,
	}
}

func (app *clientApplication) CreateClient(ctx context.Context, req *client.CreateClientRequest) (*client.CreateClientResponse, error) {
	clientEntity := convert.ClientEntityFromCreateClientRequest(req)
	if clientEntity.Secret == "" {
		return nil, errors.New("secret is required")
	}

	scopeIDs := req.GetScopes()
	clientID, err := app.clientBiz.CreateClient(ctx, clientEntity, scopeIDs)
	if err != nil {
		return nil, err
	}

	return &client.CreateClientResponse{
		ClientId: clientID,
		BaseResp: rpc.Success(),
	}, nil
}

func (app *clientApplication) ValidateClient(ctx context.Context, req *client.ValidateClientRequest) (*client.ValidateClientResponse, error) {
	clientID := req.GetClientId()
	if clientID == 0 {
		return nil, errors.New("client_id is required")
	}

	isOK, reason, err := app.clientBiz.ValidateClient(ctx, clientID, req.GetScope())
	if err != nil {
		return nil, err
	}

	return &client.ValidateClientResponse{
		IsOk:     isOK,
		Reason:   reason,
		BaseResp: rpc.Success(),
	}, nil
}
