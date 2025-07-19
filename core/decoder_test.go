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
