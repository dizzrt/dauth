package sp

import (
	"context"

	sp_api "github.com/dizzrt/dauth/api/gen/sp"
	"github.com/dizzrt/dauth/internal/domain/sp/entity"
	"github.com/dizzrt/dauth/internal/domain/sp/repo"
	"github.com/dizzrt/dauth/internal/infra/foundation"
	model "github.com/dizzrt/dauth/internal/infra/repo/model/sp"
)

var _ repo.ServiceProviderRepo = (*ServiceProviderRepoImpl)(nil)

type ServiceProviderRepoImpl struct {
	*foundation.BaseDB
}

func NewServiceProviderRepoImpl(base *foundation.BaseDB) repo.ServiceProviderRepo {
	return &ServiceProviderRepoImpl{
		BaseDB: base,
	}
}

func (impl *ServiceProviderRepoImpl) Create(ctx context.Context, sp *entity.ServiceProvider) (*entity.ServiceProvider, error) {
	m := &model.ServiceProvider{
		Name:        sp.Name,
		Description: sp.Description,
		Secret:      sp.Secret,
		RedirectURI: sp.RedirectURI,
		Status:      uint(sp_api.ServiceProvider_Status(sp.Status)),
	}

	db := impl.WithContext(ctx)
	if err := db.Create(&m).Error; err != nil {
		return nil, impl.WrapError(err)
	}

	return m.ToEntity(), nil
}

func (impl *ServiceProviderRepoImpl) Get(ctx context.Context, id uint32) (*entity.ServiceProvider, error) {
	var m *model.ServiceProvider
	db := impl.WithContext(ctx)
	if err := db.Where("id = ?", id).First(&m).Error; err != nil {
		return nil, impl.WrapError(err)
	}

	return m.ToEntity(), nil
}

func (impl *ServiceProviderRepoImpl) List(ctx context.Context, size, offset uint32) ([]*entity.ServiceProvider, uint32, error) {
	var ms []*model.ServiceProvider
	db := impl.WithContext(ctx)
	if err := db.Order("id desc").Limit(int(size)).Offset(int(offset)).Find(&ms).Error; err != nil {
		return nil, 0, impl.WrapError(err)
	}

	var total int64
	if err := db.Model(&model.ServiceProvider{}).Count(&total).Error; err != nil {
		return nil, 0, impl.WrapError(err)
	}

	respList := make([]*entity.ServiceProvider, 0, len(ms))
	for _, m := range ms {
		respList = append(respList, m.ToEntity())
	}

	return respList, uint32(total), nil
}
