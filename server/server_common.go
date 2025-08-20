package server

import (
	"godis/core"
	"godis/core/structs"
	"io"
	"log"
	"strings"
)

func readClient(conn io.ReadWriter) (*structs.RedisCommands, error) {
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

func writeClient(conn io.ReadWriter, read any) {
	encoded, err := core.Encode(read, true)
	if err != nil {
		log.Print("[ERROR] WRITING CLIENT ", conn, err)
		out, _ := core.Encode(err, true)
		conn.Write(out)
	}
	conn.Write(encoded)
}
