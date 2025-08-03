package core

import (
	"errors"
	"strconv"
)

func DecodeCommands(in []byte) ([]string, error) {
	val, _, err := Decode(in)
	if err != nil {
		return nil, err
	}

	commands := make([]string, len(val.([]any)))
	for i, v := range val.([]any) {
		commands[i] = v.(string)
	}

	return commands, err
}

func Decode(in []byte) (any, int, error) {
	if len(in) == 0 {
		return nil, 0, errors.New("empty input")
	}
	return decode(in)
}

// https://redis.io/docs/latest/develop/reference/protocol-spec/
func decode(in []byte) (any, int, error) {
	switch string(in[0]) {
	case "+":
		return decodeSimpleString(in)
	case "$":
		return decodeBulkString(in)
	case "-":
		return decodeSimpleError(in)
	case ":":
		return decodeInteger(in)
	case "*":
		return decodeArray(in)
		/*	case "_":
			return decodeNull(in)*/
	}
	return nil, 0, errors.New("invalid input")
}

/*func decodeNull(in []byte) (any, error) {

}*/

func decodeSimpleString(in []byte) (any, int, error) {
	idx := 1
	idx = getCRLFIdx(in, idx)

	return string(in[1:idx]), idx + 2, nil
}

func decodeBulkString(in []byte) (any, int, error) {
	idx := getCRLFIdx(in, 1)
	end, err := strconv.Atoi(string(in[1:idx]))
	if err != nil {
		return nil, 0, errors.New("invalid length")
	}
	idx += 2 //\r\n
	end += idx
	return string(in[idx:end]), end + 2, nil
}

func decodeSimpleError(in []byte) (any, int, error) {
	return decodeSimpleString(in)
}

func decodeInteger(in []byte) (any, int, error) {
	idx := 1
	idx = getCRLFIdx(in, idx)
	res, err := strconv.Atoi(string(in[1:idx]))
	if err != nil {
		return nil, idx + 2, errors.New("invalid integer")
	}
	return res, idx + 2, nil
}

func decodeArray(in []byte) (any, int, error) {
	arrayLength, idx, _ := decodeInteger(in)

	array := make([]any, arrayLength.(int))
	for i := range array {
		res, nextIdx, err := Decode(in[idx:])
		if err != nil {
			return nil, 0, err
		}
		array[i] = res
		idx += nextIdx
	}
	return array, idx, nil
}

func getCRLFIdx(in []byte, idx int) int {
	for idx < len(in) {
		if in[idx] == '\r' {
			break
		}
		idx++
	}
	return idx
}
