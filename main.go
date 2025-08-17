package main

import (
	"flag"
	"godis/server"
)

func initFlags() {
	flag.IntVar(&server.Port, "p", 8080, "port to listen on")
	flag.StringVar(&server.Host, "h", "127.0.0.1", "host to listen on")
	flag.BoolVar(&server.SyncFlag, "s", false, "true to start sync server")
	flag.IntVar(&server.MAX_CLIENTS, "m", 10000, "max clients connections to accept")

	flag.Parse()
}

func main() {
	initFlags()
	if server.SyncFlag {
		server.HandleSync()
	} else {
		server.HandleAsync()
	}
}
