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
	Create(ctx context.Context, req *client.CreateRequest) (*client.CreateResponse, error)
	Validate(ctx context.Context, req *client.ValidateRequest) (*client.ValidateResponse, error)
}

type clientApplication struct {
	clientBiz biz.ClientBiz
}

func NewClientApplication(clientBiz biz.ClientBiz) ClientApplication {
	return &clientApplication{
		clientBiz: clientBiz,
	}
}

func (app *clientApplication) Create(ctx context.Context, req *client.CreateRequest) (*client.CreateResponse, error) {
	clientEntity := convert.ClientEntityFromCreateRequest(req)
	if clientEntity.Secret == "" {
		return nil, errors.New("secret is required")
	}

	scopeIDs := req.GetScopes()
	clientID, err := app.clientBiz.Create(ctx, clientEntity, scopeIDs)
	if err != nil {
		return nil, err
	}

	return &client.CreateResponse{
		ClientId: clientID,
		BaseResp: rpc.Success(),
	}, nil
}

func (app *clientApplication) Validate(ctx context.Context, req *client.ValidateRequest) (*client.ValidateResponse, error) {
	clientID := req.GetClientId()
	if clientID == 0 {
		return nil, errors.New("client_id is required")
	}

	isOK, reason, err := app.clientBiz.Validate(ctx, clientID, req.GetScope())
	if err != nil {
		return nil, err
	}

	return &client.ValidateResponse{
		IsOk:     isOK,
		Reason:   reason,
		BaseResp: rpc.Success(),
	}, nil
}
