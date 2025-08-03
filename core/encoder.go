package core

import (
	"errors"
	"fmt"
)

const SIMPLE_STRING_ENCODING = "+%s\r\n"
const INTEGER_STRING_ENCODING = ":%d\r\n"
const BULK_STRING_ENCODING = "$%d\r\n%s\r\n"

// Encode strings to RESP Encoding+
func Encode(input any, isSimple bool) ([]byte, error) {
	switch in := input.(type) {
	case string:
		if isSimple {
			return []byte(fmt.Sprintf(SIMPLE_STRING_ENCODING, in)), nil
		}
		return []byte(fmt.Sprintf(BULK_STRING_ENCODING, len(in), in)), nil
	case int:
		return []byte(fmt.Sprintf(INTEGER_STRING_ENCODING, in)), nil
	}
	return nil, errors.New(fmt.Sprint("Invalid Input"))
}
