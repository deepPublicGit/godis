package core

import (
	"errors"
	"godis/core/structs"
	"godis/core/utils"
	"strconv"
	"time"
)

func Eval(commands *structs.RedisCommands) (any, error) {
	switch commands.Cmd {
	case "PING":
		return evalPING(commands.Args)
	case "GET":
		return evalGET(commands.Args)
	case "SET":
		return evalSET(commands.Args)
	case "INCR":
		return evalADD(commands.Args, 1)
	case "DECR":
		return evalADD(commands.Args, -1)
	case "DEL":
		return evalDEL(commands.Args)
	case "TTL":
		return evalTTL(commands.Args)
	case "EXPIRE":
		return evalEXPIRE(commands.Args)
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

func evalGET(args []string) (any, error) {
	if len(args) <= 1 {
		return "", errors.New("invalid number of arguments for GET")
	}
	key := args[0]
	redisObject, ok := structs.Get(key)
	if ok {
		return redisObject, nil
	}
	return "-1", nil
}

func evalSET(args []string) (string, error) {
	if len(args) <= 1 {
		return "", errors.New("invalid number of arguments for SET")
	}

	key, val, expMs := args[0], args[1], int64(-1)
	for i := 2; i < len(args); i++ {
		switch args[i] {
		case "EX":
			i++
			if i >= len(args) {
				return "", errors.New("expiry value missing for EX")
			}

			expS, err := strconv.ParseInt(args[i+1], 10, 64)
			if err != nil {
				return "", errors.New("invalid or out of range expiry value")
			}
			expMs = expS * 1000
		default:
			return "", errors.New("invalid argument")
		}
	}
	structs.Set(key, structs.NewRedisObject(val, expMs))
	return "OK", nil
}

// evalINCR op indicates value to add or subtract
func evalADD(args []string, op int) (string, error) {
	if len(args) <= 1 {
		return "", errors.New("invalid number of arguments for INCR")
	}
	key := args[1]
	redisObject, ok := structs.Get(key)
	if ok {
		val := redisObject.Value.(string)
		incr, err := strconv.Atoi(val)
		if err != nil {
			return evalSET([]string{key, "1"})
		}
		return evalSET([]string{key, strconv.Itoa(incr + op)})
	}
	return evalSET([]string{key, "1"})
}

func evalDEL(args []string) (any, error) {
	if len(args) < 1 {
		return "", errors.New("invalid number of keys to delete for DEL")
	}
	keysDeleted := 0
	for _, key := range args {
		if structs.Del(key) {
			keysDeleted++
		}
	}
	return keysDeleted, nil
}

func evalTTL(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("invalid number of arguments for TTL")
	}
	key := args[0]
	redisObject, ok := structs.Get(key)
	if ok {
		if redisObject.ExpiresAt == -1 {
			return "-1", nil
		}
		ttl := (redisObject.ExpiresAt - time.Now().UnixMilli()) / 1000
		if ttl > 0 {
			return strconv.FormatInt(ttl, 10), nil
		}
	}
	return "-2", nil
}

func evalEXPIRE(args []string) (string, error) {
	if len(args) <= 2 {
		return "", errors.New("invalid number of arguments for EXPIRE")
	}
	key, expS := args[1], args[2]
	exp, err := strconv.ParseInt(expS, 10, 64)

	if err != nil || exp < -1 {
		return "-1", errors.New("invalid or out of range expiry value")
	}

	redisObject, ok := structs.Get(key)
	if ok {
		redisObject.ExpiresAt = utils.GetExpiryInUnixMs(exp * 1000)
		return "1", nil
	}

	return "0", nil
}
