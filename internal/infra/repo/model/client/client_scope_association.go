package client

import (
	"github.com/dizzrt/dauth/internal/domain/client/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ schema.Tabler = (*ClientScopeAssociation)(nil)

type ClientScopeAssociation struct {
	gorm.Model
	ClientID uint32 `gorm:"column:client_id"`
	ScopeID  uint32 `gorm:"column:scope_id"`
}

func (c *ClientScopeAssociation) TableName() string {
	return "client_client_scope_associations"
}

func (c *ClientScopeAssociation) ToEntity() *entity.ClientScopeAssociation {
	return &entity.ClientScopeAssociation{
		ID:        uint32(c.ID),
		ClientID:  c.ClientID,
		ScopeID:   c.ScopeID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: c.DeletedAt,
	}
}
