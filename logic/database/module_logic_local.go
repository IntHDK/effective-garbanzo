package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseModule_Local struct {
	db     *gorm.DB
	logger *log.Logger
}

func (module DatabaseModule_Local) connect(DSN string, DBType string, Migrate bool) (result bool, err error) {
	switch DBType {
	case DBTYPE_MYSQL:
		module.db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
		if err == nil && module.db != nil {
			module.logger.Printf("DB connected, DBType = %s, DSN = %s\n", DBType, DSN)
			result = true
		}
	default:
		result = false
		return
	}
	if err != nil {
		result = false
		return
	}

	if Migrate {
		errmigrate := module.db.AutoMigrate(&ModelPost{})
		if errmigrate != nil {
			module.logger.Fatalf("DB migraion error : %s\n", errmigrate.Error())
			return
		}
	}

	return
}

func (module DatabaseModule_Local) disconnect() (result bool, err error) {
	module.db = nil

	result = true
	err = nil
	return
}

func (module DatabaseModule_Local) IsReady() (ready bool) {
	return (module.db != nil)
}

func NewDatabaseModule_Local(DSN string, DBType string, Migrate bool, Logger *log.Logger) (res DatabaseModule, err error) {
	dbmodule := DatabaseModule_Local{
		logger: Logger,
	}
	resmodule, err := dbmodule.connect(DSN, DBType, Migrate)
	if resmodule {
		res = dbmodule
	} else {
		res = nil
	}
	return
}
