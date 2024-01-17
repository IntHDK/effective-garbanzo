package logic

import (
	"effective-garbanzo/logic/database"
	"log"
)

type LogicModule struct {
	database database.DatabaseModule
	Logger   *log.Logger
}
type LogicModuleConfiguration struct {
	Database database.DatabaseModule
	Logger   *log.Logger
}

func NewLogicModule(configuration LogicModuleConfiguration) (module *LogicModule) {
	module = &LogicModule{
		database: configuration.Database,
		Logger:   configuration.Logger,
	}
	return
}
