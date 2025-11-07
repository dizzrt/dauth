package model

import (
	"time"

	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email         string `gorm:"unique"`
	Username      string `gorm:"not null"`
	Password      string `gorm:"not null"`
	Status        uint   `gorm:"not null"`
	LastLoginTime time.Time
}

func ToEntity(m *User) (*entity.User, error) {
	var e entity.User
	// 配置字段映射（Detail -> Info，Username -> Name）
	config := &mapstructure.DecoderConfig{
		Result:  &e,
		TagName: "mapstructure", // 用标签指定映射关系
		// IgnoreUnknownFields: true,           // 忽略 entity 中不存在的字段（如 model 中的存储字段）
		// IgnoreUntaggedFields: true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}
	// 将 model 转为 map 后解码（支持嵌套）
	return &e, decoder.Decode(m)
}
