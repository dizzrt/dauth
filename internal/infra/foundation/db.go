package foundation

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dizzrt/dauth/api/gen/errdef"
	"github.com/dizzrt/dauth/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseDB struct {
	db *gorm.DB
}

func NewBaseDB(ac *conf.AppConfig) *BaseDB {
	return &BaseDB{
		db: newDB(ac),
	}
}

func (base *BaseDB) WrapError(err error) error {
	if err == nil {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errdef.RecordNotFound().WithCause(err)
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return errdef.DuplicatedKey().WithCause(err)
	}

	return err
}

func (base *BaseDB) WithContext(ctx context.Context) *gorm.DB {
	return base.db.WithContext(ctx)
}

func buildDSN(ac *conf.AppConfig) string {
	user := ac.DB.User
	password := ac.DB.Password
	database := ac.DB.Database
	addr := ac.DB.Addr

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, addr, database)
}

func newDB(ac *conf.AppConfig) *gorm.DB {
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
