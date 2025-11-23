package client

import (
	"github.com/dizzrt/dauth/api/gen/client"
	"github.com/dizzrt/dauth/internal/domain/client/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ schema.Tabler = (*Client)(nil)

type Client struct {
	gorm.Model
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Secret      string `gorm:"column:secret"`
	RedirectURI string `gorm:"column:redirect_uri"`
	Status      uint   `gorm:"column:status"`
}

func (c *Client) TableName() string {
	return "client_clients"
}

func (c *Client) ToEntity() *entity.Client {
	return &entity.Client{
		ID:          uint32(c.ID),
		Name:        c.Name,
		Description: c.Description,
		Secret:      c.Secret,
		RedirectURI: c.RedirectURI,
		Status:      client.Client_Status(c.Status),
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		DeletedAt:   c.DeletedAt,
	}
}
