package main

import (
	"effective-garbanzo/logic"
	"effective-garbanzo/webserver"
	"flag"
	"log"
)

var addrflag = flag.String("addr", "localhost:28774", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	var addr = "localhost:28774"
	if addrflag != nil {
		addr = *addrflag
	}

	logicmodule := logic.NewLogicModule(logic.LogicModuleConfiguration{
		Database: nil,
		Logger:   log.Default(),
	})

	ws := webserver.NewWebAllServer(webserver.WebAllServerConfiguration{
		ListenAt:    addr,
		Logger:      log.Default(),
		LogicModule: logicmodule,
	})
	log.Fatal(ws.HttpServerStart())
}
