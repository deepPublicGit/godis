package core

import "testing"

func TestDecodeSimpleString(t *testing.T) {
	cases := map[string]string{
		"+PING\r\n": "PING",
	}
	for in, out := range cases {
		actual, _ := Decode([]byte(in))
		if actual != out {
			t.Errorf("Expected %q, Actual %q", out, actual)
		}
	}
}

func TestDecodeBulkString(t *testing.T) {
	cases := map[string]string{
		"$4\r\nPING\r\n": "PING",
		"$0\r\n\r\n":     "",
	}
	for in, out := range cases {
		actual, _ := Decode([]byte(in))
		if actual != out {
			t.Errorf("Expected %q, Actual %q", out, actual)
		}
	}
}

func TestDecodeArray(t *testing.T) {
	cases := map[string][]interface{}{
		"*1\r\n+PING\r\n":                      {"PING"},             //simple string
		"*1\r\n$4\r\nPING\r\n":                 {"PING"},             //bulk string
		"*2\r\n$4\r\nPING\r\n+PONG\r\n":        {"PING", "PONG"},     //simple & bulk string
		"*3\r\n$4\r\nPING\r\n+PONG\r\n:10\r\n": {"PING", "PONG", 10}, //mix
		"*5\r\n$4\r\nPING\r\n+PONG\r\n:10\r\n*0\r\n*2\r\n$4\r\nPING2\r\n+PONG2\r\n": {"PING", "PONG", 10, []interface{}, {"PING2, PONG2"}}, //mix with array
		"*0\r\n": {},
	}
	for in, out := range cases {
		actual, _ := Decode([]byte(in))
		arr := actual.([]interface{})
		if len(arr) != len(out) {
			t.Errorf("Length Expected %q, Actual %q", len(out), len(arr))
		}

		for i := range arr {
			if arr[i] != out[i] {
				t.Errorf("Array Expected %q, Actual %q", out[i], arr[i])
			}
		}
	}
}

func TestDecodeInteger(t *testing.T) {
	cases := map[string]int{
		":101\r\n": 101,
	}
	for in, out := range cases {
		actual, _ := Decode([]byte(in))
		if actual != out {
			t.Errorf("Expected %q, Actual %q", out, actual)
		}
	}
}

func TestDecodeZeroLength(t *testing.T) {
	cases := map[string]string{
		"": "empty input",
	}
	for in, out := range cases {
		actual, err := Decode([]byte(in))
		if actual != nil {
			t.Errorf("Expected Error, Actual %q", actual)
		}
		if err == nil {
			t.Errorf("Expected %q, Actual %q", out, actual)
		} else {
			if err.Error() != out {
				t.Errorf("Expected %q, Actual %q", out, err.Error())
			}
		}
	}
}
