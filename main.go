package main

import (
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

	ws := webserver.NewWebAllServer(webserver.WebAllServerConfiguration{
		ListenAt: addr,
		Logger:   log.Default(),
	})
	log.Fatal(ws.HttpServerStart())
}
