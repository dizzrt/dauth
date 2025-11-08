package common

import (
	"context"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseDB struct {
	db *gorm.DB
}

func NewBaseDB() *BaseDB {
	return &BaseDB{
		db: newDB(),
	}
}

func (base *BaseDB) WithContext(ctx context.Context) *gorm.DB {
	return base.db.WithContext(ctx)
}

func buildDSN() string {
	// TODO build from conf
	return "root:u6xYzLEu4xZQg2jPJHMk@tcp(192.168.124.10:3306)/dauth?charset=utf8mb4&parseTime=True&loc=Local"
}

func newDB() *gorm.DB {
	dsn := buildDSN()
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
