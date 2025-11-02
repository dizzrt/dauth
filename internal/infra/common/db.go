package common

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	dsn := buildDSN()
	db, err = gorm.Open(mysql.New(mysql.Config{
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
}

func buildDSN() string {
	return "root:u6xYzLEu4xZQg2jPJHMk@tcp(192.168.124.10:3306)/dauth?charset=utf8mb4&parseTime=True&loc=Local"
}

func DB() *gorm.DB {
	return db
}
