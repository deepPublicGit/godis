package core

import (
	"errors"
	"godis/core/structs"
)

func Eval(commands *structs.RedisCommands) (string, error) {
	switch commands.Cmd {
	case "PING":
		return evalPING(commands.Args)
	case "GET":
		return evalGET(commands.Args)
	case "SET":
		return evalSET(commands.Args)
	case "INCR":
		return evalINCR(commands.Args)
	case "DEL":
		return evalDEL(commands.Args)
	case "TTL":
		return evalTTL(commands.Args)
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

func evalGET(args []string) (string, error) {
	if len(args) > 1 {
		return "", errors.New("invalid number of arguments")
	}
	if len(args) == 0 {
		return "PONG", nil
	} else {
		return "PONG " + args[0], nil
	}
}

func evalSET(args []string) (string, error) {
	if len(args) > 1 {
		return "", errors.New("invalid number of arguments")
	}
	if len(args) == 0 {
		return "PONG", nil
	} else {
		return "PONG " + args[0], nil
	}
}

func evalINCR(args []string) (string, error) {
	if len(args) > 1 {
		return "", errors.New("invalid number of arguments")
	}
	if len(args) == 0 {
		return "PONG", nil
	} else {
		return "PONG " + args[0], nil
	}
}

func evalDEL(args []string) (string, error) {
	if len(args) > 1 {
		return "", errors.New("invalid number of arguments")
	}
	if len(args) == 0 {
		return "PONG", nil
	} else {
		return "PONG " + args[0], nil
	}
}

func evalTTL(args []string) (string, error) {
	if len(args) > 1 {
		return "", errors.New("invalid number of arguments")
	}
	if len(args) == 0 {
		return "PONG", nil
	} else {
		return "PONG " + args[0], nil
	}
}
