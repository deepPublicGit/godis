package core

import (
	"errors"
	"strconv"
)

func Decode(in []byte) (interface{}, error) {
	if len(in) == 0 {
		return nil, errors.New("empty input")
	}
	out, err := decode(in)
	return out, err
}

// https://redis.io/docs/latest/develop/reference/protocol-spec/
func decode(in []byte) (interface{}, error) {
	switch string(in[0]) {
	case "+":
		return decodeSimpleString(in)
	case "-":
		return decodeSimpleError(in)
	case ":":
		return decodeInteger(in)
	case "$":
		return decodeBulkString(in)
		/*	case "*":
				return decodeArray(in)
			case "_":
				return decodeNull(in)*/
	}
	return nil, errors.New("invalid input")
}

/*func decodeNull(in []byte) (interface{}, error) {

}

func decodeArray(in []byte) (interface{}, error) {

}*/

func decodeBulkString(in []byte) (interface{}, error) {
	idx := getCRLFIdx(in, 1)
	end, err := strconv.Atoi(string(in[1:idx]))
	if err != nil {
		return nil, errors.New("invalid length")
	}
	idx += 2 //\r\n
	end += idx
	return string(in[idx:end]), nil
}

func decodeInteger(in []byte) (interface{}, error) {
	idx := 1
	idx = getCRLFIdx(in, idx)
	res, err := strconv.Atoi(string(in[1:idx]))
	if err != nil {
		return nil, errors.New("invalid integer")
	}
	return res, nil
}

func decodeSimpleError(in []byte) (interface{}, error) {
	idx := 1
	idx = getCRLFIdx(in, idx)

	return nil, errors.New(string(in[1:idx]))
}

func decodeSimpleString(in []byte) (interface{}, error) {
	idx := 1
	idx = getCRLFIdx(in, idx)

	return string(in[1:idx]), nil
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
