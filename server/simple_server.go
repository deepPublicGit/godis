package server

import (
	"godis/core"
	"log"
	"net"
	"strings"
)

func Handle(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
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
				err := writeConnection(conn, err.Error())
				if err != nil {
					log.Print("[ERROR] RESPONDING CLIENT ", conn.RemoteAddr(), err)
				}
			}
			log.Print("[DEBUG] CLIENT INPUT: ", output)
			if err = writeConnection(conn, output); err != nil {
				log.Print("[ERROR] RESPONDING CLIENT ", conn.RemoteAddr(), output)
			}
		}
	}
}

func readCommands(conn net.Conn) (*core.RedisCommands, error) {
	buffer := make([]byte, 256)
	size, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	commands, err := core.DecodeCommands(buffer[:size])
	if err != nil {
		return nil, err
	}

	return &core.RedisCommands{
		Cmd:  strings.ToUpper(commands[0]),
		Args: commands[1:],
	}, nil
}

func writeConnection(conn net.Conn, read string) error {
	_, err := conn.Write([]byte("SERVER: " + read))
	if err != nil {
		return err
	}
	return nil
}
