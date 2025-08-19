package core

import (
	"errors"
	"godis/core/structs"
)

func Eval(commands *structs.RedisCommands) (string, error) {
	switch commands.Cmd {
	case "PING":
		return evalPING(commands.Args)
	default:
		return "", errors.New("invalid command")
	}
}

func evalPING(args []string) (string, error) {
	if len(args) > 1 {
		return "", errors.New("invalid number of arguments")
	}
	if len(args) == 0 {
		return "PONG", nil
	} else {
		return "PONG " + args[0], nil
	}
}
