package common

import (
	"context"
	"fmt"
	"time"

	"github.com/dizzrt/dauth/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseDB struct {
	db *gorm.DB
}

func NewBaseDB(bootstrap *conf.Bootstrap) *BaseDB {
	return &BaseDB{
		db: newDB(bootstrap),
	}
}

func (base *BaseDB) WithContext(ctx context.Context) *gorm.DB {
	return base.db.WithContext(ctx)
}

func buildDSN(bootstrap *conf.Bootstrap) string {
	user := bootstrap.DB.User
	password := bootstrap.DB.Password
	database := bootstrap.DB.Database
	addr := bootstrap.DB.Addr

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, addr, database)
}

func newDB(bootstrap *conf.Bootstrap) *gorm.DB {
	dsn := buildDSN(bootstrap)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
