package server

import (
	"fmt"
	"godis/core"
	"godis/core/structs"
	"io"
	"log"
	"net"
	"strings"
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

func readCommands(conn io.ReadWriter) (*structs.RedisCommands, error) {
	buffer := make([]byte, 256)
	size, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	commands, err := core.DecodeCommands(buffer[:size])
	if err != nil {
		return nil, err
	}

	return &structs.RedisCommands{
		Cmd:  strings.ToUpper(commands[0]),
		Args: commands[1:],
	}, nil
}

func writeConnection(conn io.ReadWriter, read any) {
	encoded, err := core.Encode(read, true)
	if err != nil {
		log.Print("[ERROR] WRITING CLIENT ", conn, err)
		out, _ := core.Encode(err, true)
		conn.Write(out)
	}
	conn.Write(encoded)
}
