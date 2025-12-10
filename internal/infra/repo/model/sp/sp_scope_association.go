package sp

import (
	"github.com/dizzrt/dauth/internal/domain/sp/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ schema.Tabler = (*SPScopeAssociation)(nil)

type SPScopeAssociation struct {
	gorm.Model
	SPID    uint32 `gorm:"column:sp_id"`
	ScopeID uint32 `gorm:"column:scope_id"`
}

func (c *SPScopeAssociation) TableName() string {
	return "sp_sp_scope_associations"
}

func (c *SPScopeAssociation) ToEntity() *entity.SPScopeAssociation {
	return &entity.SPScopeAssociation{
		ID:        uint32(c.ID),
		SPID:      c.SPID,
		ScopeID:   c.ScopeID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: c.DeletedAt,
	}
}
