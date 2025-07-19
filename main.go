package main

import (
	"flag"
	"fmt"
	"godis/server"
	"log"
	"net"
)

var port int
var host string

func initFlags() {
	flag.IntVar(&port, "p", 8080, "port to listen on")
	flag.StringVar(&host, "h", "127.0.0.1", "host to listen on")
	flag.Parse()
}

func main() {
	initFlags()
	log.Printf("Listening on %s:%d", host, port)
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}
	server.Handle(listen)
}

// Start server
