package client

import (
	"context"

	client_api "github.com/dizzrt/dauth/api/gen/client"
	"github.com/dizzrt/dauth/internal/domain/client/entity"
	"github.com/dizzrt/dauth/internal/domain/client/repo"
	"github.com/dizzrt/dauth/internal/infra/foundation"
	model "github.com/dizzrt/dauth/internal/infra/repo/model/client"
)

var _ repo.ClientRepo = (*ClientRepoImpl)(nil)

type ClientRepoImpl struct {
	*foundation.BaseDB
}

func NewClientRepoImpl(base *foundation.BaseDB) repo.ClientRepo {
	return &ClientRepoImpl{
		BaseDB: base,
	}
}

func (impl *ClientRepoImpl) Create(ctx context.Context, client *entity.Client) (*entity.Client, error) {
	m := &model.Client{
		Name:        client.Name,
		Description: client.Description,
		Secret:      client.Secret,
		RedirectURI: client.RedirectURI,
		Status:      uint(client_api.Client_ACTIVE),
	}

	db := impl.WithContext(ctx)
	if err := db.Create(&m).Error; err != nil {
		return nil, impl.WrapError(err)
	}

	return m.ToEntity(), nil
}

func (impl *ClientRepoImpl) Get(ctx context.Context, id uint32) (*entity.Client, error) {
	var m *model.Client
	db := impl.WithContext(ctx)
	if err := db.Where("id = ?", id).First(&m).Error; err != nil {
		return nil, impl.WrapError(err)
	}

	return m.ToEntity(), nil
}
