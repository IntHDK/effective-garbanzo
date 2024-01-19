package main

import (
	"effective-garbanzo/logic"
	"effective-garbanzo/logic/database"
	"effective-garbanzo/webserver"
	"flag"
	"log"
)

var addrflag = flag.String("addr", "localhost:28774", "http service address")
var dbflag = flag.String("db", "gorm:gorm@tcp(172.30.176.1:3306)/effective-garbanzo?charset=utf8mb4&parseTime=True&loc=Local", "(temp) database")

func main() {
	flag.Parse()
	log.SetFlags(0)
	var addr = "localhost:28774"
	if addrflag != nil {
		addr = *addrflag
	}
	var param_db = "gorm:gorm@tcp(172.30.176.1:3306)/effective-garbanzo?charset=utf8mb4&parseTime=True&loc=Local"
	if dbflag != nil {
		param_db = *dbflag
	}

	dbmodule, err := database.NewDatabaseModule_Local(param_db, database.DBTYPE_MYSQL, true, log.Default())
	if err != nil {
		//TODO: db 접근실패시 memory 또는 retry
		log.Fatalf("db error : %v\n", err)
		return
	}

	logicmodule := logic.NewLogicModule(logic.LogicModuleConfiguration{
		Database: dbmodule,
		Logger:   log.Default(),
	})

	ws := webserver.NewWebAllServer(webserver.WebAllServerConfiguration{
		ListenAt:    addr,
		Logger:      log.Default(),
		LogicModule: logicmodule,
	})
	log.Fatal(ws.HttpServerStart())
}
