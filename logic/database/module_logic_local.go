package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseModule_Local struct {
	db *gorm.DB
}

func (module DatabaseModule_Local) Connect(DSN string, DBType string, Migrate bool) (result bool, err error) {
	switch DBType {
	case DBTYPE_MYSQL:
		module.db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
		if err == nil && module.db != nil {
			result = true
		}
	default:
		result = false
		return
	}

	if Migrate {
		module.db.AutoMigrate(&ModelPost{})
	}

	return
}

func (module DatabaseModule_Local) Disconnect() (result bool, err error) {
	module.db = nil

	result = true
	err = nil
	return
}

func (module DatabaseModule_Local) IsReady() (ready bool) {
	return (module.db != nil)
}

func NewDatabaseModule_Local(DSN string, DBType string, Migrate bool) (res DatabaseModule, err error) {
	dbmodule := DatabaseModule_Local{}
	resmodule, err := dbmodule.Connect(DSN, DBType, Migrate)
	if resmodule {
		res = dbmodule
	} else {
		res = nil
	}
	return
}
