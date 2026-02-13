package foundation

import (
	"fmt"
	"time"

	"github.com/dizzrt/dauth/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func buildDSN(ac *conf.AppConfig) string {
	user := ac.DB.User
	password := ac.DB.Password
	database := ac.DB.Database
	addr := ac.DB.Addr

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, addr, database)
}

func NewDB(ac *conf.AppConfig) *gorm.DB {
	dsn := buildDSN(ac)
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
