package server

import (
	"log"
	"net"
)

func Handle(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		log.Print("[STARTED] PROCESSING CLIENT ", conn.RemoteAddr())

		for {
			read, err := readConnection(conn)
			if err != nil {
				conn.Close()
				log.Print("[ERROR] READING CLIENT ", conn.RemoteAddr())

				break
			}
			log.Print("[DEBUG] CLIENT INPUT: ", read)
			if err = writeConnection(conn, read); err != nil {
				log.Print("[ERROR] RESPONDING CLIENT ", conn.RemoteAddr(), read)
			}
		}
	}
}

func readConnection(conn net.Conn) (string, error) {
	buffer := make([]byte, 256)
	size, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:size]), nil
}

func writeConnection(conn net.Conn, read string) error {
	_, err := conn.Write([]byte("SERVER: " + read))
	if err != nil {
		return err
	}
	return nil
}
