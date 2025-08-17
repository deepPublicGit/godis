package server

import (
	"godis/core"
	"log"
	"syscall"
)

func HandleAsync() {

	log.Printf("Listening async on %s:%d", Host, Port)

	events := make([]syscall.EpollEvent, 100)
	fd, err := syscall.Socket(syscall.AF_INET, syscall.O_NONBLOCK|syscall.SOCK_STREAM, 0)

	if err != nil {
		panic(err)
	}

	defer syscall.Close(fd)

	log.Print("[STARTED] PROCESSING CLIENT ", conn.RemoteAddr())

	for {
		commands, err := readCommands(conn)
		if err != nil {
			conn.Close()
			log.Print("[ERROR] READING CLIENT ", conn.RemoteAddr(), err)

			break
		}
		output, err := core.Eval(commands)
		if err != nil {
			writeConnection(conn, err)
		}
		log.Print("[DEBUG] CLIENT INPUT: ", output)
		writeConnection(conn, output)
	}
}
