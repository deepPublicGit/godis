package main

import (
	"flag"
	"godis/server"
)

func initFlags() {
	flag.IntVar(&server.Port, "p", 8080, "Port to listen on")
	flag.StringVar(&server.Host, "h", "127.0.0.1", "host to listen on")
	flag.IntVar(&server.SyncFlag, "s", 0, "1 to start sync server")
	flag.Parse()
}

func main() {
	initFlags()
	if server.SyncFlag == 1 {
		server.HandleSync()
	} else {
		server.HandleAsync()
	}
}

// Start server
