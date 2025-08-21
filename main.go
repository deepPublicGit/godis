package main

import (
	"flag"
	"godis/config"
	"godis/server"
)

func initFlags() {
	flag.IntVar(&config.Port, "p", 8080, "port to listen on")
	flag.StringVar(&config.Host, "h", "127.0.0.1", "host to listen on")
	flag.BoolVar(&config.SyncFlag, "s", false, "true to start sync server")
	flag.IntVar(&config.MAX_CLIENTS, "m", 10000, "max clients connections to accept")
	flag.IntVar(&config.ExpiryCron, "e", 5, "number of seconds between expiry cycles")
	flag.IntVar(&config.ExpirySample, "d", 25, "sample percentage to delete from expiryLimit")
	flag.IntVar(&config.ExpiryLimit, "l", 20, "max number of keys to be deleted")
	flag.IntVar(&config.MaxMemory, "M", 1024, "max memory in MB to use before keys are evicted")
	flag.StringVar(&config.EvictionStrategy, "E", "lru", "eviction strategy to use")

	flag.Parse()
}

func main() {
	initFlags()
	if config.SyncFlag {
		server.HandleSync()
	} else {
		server.HandleAsync()
	}
}
