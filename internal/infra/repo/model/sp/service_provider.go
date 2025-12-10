package sp

import (
	"github.com/dizzrt/dauth/api/gen/sp"
	"github.com/dizzrt/dauth/internal/domain/sp/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ schema.Tabler = (*ServiceProvider)(nil)

type ServiceProvider struct {
	gorm.Model
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Secret      string `gorm:"column:secret"`
	RedirectURI string `gorm:"column:redirect_uri"`
	Status      uint   `gorm:"column:status"`
}

func (c *ServiceProvider) TableName() string {
	return "sp_service_providers"
}

func (c *ServiceProvider) ToEntity() *entity.ServiceProvider {
	return &entity.ServiceProvider{
		ID:          uint32(c.ID),
		Name:        c.Name,
		Description: c.Description,
		Secret:      c.Secret,
		RedirectURI: c.RedirectURI,
		Status:      sp.ServiceProvider_Status(c.Status),
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		DeletedAt:   c.DeletedAt,
	}
}
