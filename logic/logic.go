package logic

import (
	"effective-garbanzo/logic/database"
)

type LogicModule struct {
	Database database.DatabaseModule
}
