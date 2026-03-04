package utils

import (
	"errors"

	"github.com/dizzrt/dauth/api/gen/errdef"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func WrapError(err error) error {
	if err == nil {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errdef.RecordNotFound().WithCause(err)
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return errdef.DuplicatedKey().WithCause(err)
	}

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return errdef.DuplicatedKey().WithCause(err)
	}

	return err
}
