package server

import (
	"fmt"
	"godis/core"
	"log"
	"net"
)

// HandleSync For local testing only
func HandleSync() {

	log.Printf("Listening sync on %s:%d", Host, Port)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", Host, Port))

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		log.Print("[STARTED] PROCESSING CLIENT ", conn.RemoteAddr())

		for {
			commands, err := readClient(conn)
			if err != nil {
				conn.Close()
				log.Print("[ERROR] READING CLIENT ", conn.RemoteAddr(), err)

				break
			}
			output, err := core.Eval(commands)
			if err != nil {
				writeClient(conn, err)
			}
			log.Print("[DEBUG] CLIENT INPUT: ", output)
			writeClient(conn, output)
		}
	}
}
