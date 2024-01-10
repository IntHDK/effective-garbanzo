package database

import (
	"errors"

	"gorm.io/gorm"
)

var ERROR_DUPLICATED error
var ERROR_PASSWORDINCORRECT error

func init_errors() {
	ERROR_DUPLICATED = gorm.ErrDuplicatedKey
	ERROR_PASSWORDINCORRECT = errors.New(ERROR_PASSWORDHASH_INCORRECT)
}
