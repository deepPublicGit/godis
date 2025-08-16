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
				writeConnection(conn, err)
			}
			log.Print("[DEBUG] CLIENT INPUT: ", output)
			writeConnection(conn, output)
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

func writeConnection(conn net.Conn, read any) {
	encoded, err := core.Encode(read, true)
	if err != nil {
		log.Print("[ERROR] WRITING CLIENT ", conn.RemoteAddr(), err)
		out, _ := core.Encode(err, true)
		conn.Write(out)
	}
	conn.Write(encoded)
}
