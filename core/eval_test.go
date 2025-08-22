package core

import (
	"godis/core/structs"
	"testing"
)

func TestEvalPingValidCases(t *testing.T) {
	cases := map[*structs.RedisCmd]string{
		&structs.RedisCmd{Cmd: "PING", Args: nil}:            "PONG",
		&structs.RedisCmd{Cmd: "PING", Args: []string{"OK"}}: "PONG OK",
	}
	for in, out := range cases {
		actual, err := Eval(in)
		if err != nil {
			t.Fatalf("Eval(%q): %v", in, err)
		}
		if actual != out {
			t.Errorf("Expected %q, Actual %q", out, actual)
		}
	}
}

func TestEvalPingInvalidCases(t *testing.T) {
	cases := map[*structs.RedisCmd]string{
		&structs.RedisCmd{Cmd: "PONG", Args: nil}:                       "invalid command",
		&structs.RedisCmd{Cmd: "PING", Args: []string{"OKIE", "DOKIE"}}: "invalid number of arguments",
	}
	for in, out := range cases {
		actual, err := Eval(in)
		if err == nil || actual != "" {
			t.Errorf("Expected error: %s, for %s", err, in)
		}
		if err.Error() != out {
			t.Errorf("Expected %q, Actual %q", out, err.Error())
		}
	}
}

func TestEvalGet(t *testing.T) {
	cases := map[*structs.RedisCmd]string{
		&structs.RedisCmd{Cmd: "GET", Args: nil}:            "invalid number of arguments",
		&structs.RedisCmd{Cmd: "GET", Args: []string{"K1"}}: "V1",
	}
	for in, out := range cases {
		actual, err := Eval(in)
		if err != nil {
			t.Fatalf("Eval(%q): %v", in, err)
		}
		if actual != out {
			t.Errorf("Expected %q, Actual %q", out, actual)
		}
	}
}

func TestEvalSet(t *testing.T) {
	cases := map[*structs.RedisCmd]string{
		&structs.RedisCmd{Cmd: "SET", Args: nil}:                                "invalid number of arguments",
		&structs.RedisCmd{Cmd: "SET", Args: []string{"K1", "10"}}:               "OK",
		&structs.RedisCmd{Cmd: "SET", Args: []string{"K1", "10", "EX", "10"}}:   "OK",
		&structs.RedisCmd{Cmd: "SET", Args: []string{"K1", "10", "YOLO", "10"}}: "OK",
		&structs.RedisCmd{Cmd: "SET", Args: []string{"K1", "10", "EX"}}:         "expiry value missing for EX",
		&structs.RedisCmd{Cmd: "SET", Args: []string{"K1", "10", "EX", "-2"}}:   "invalid or out of range expiry value",
	}
	for in, out := range cases {
		actual, err := Eval(in)
		if err != nil {
			t.Fatalf("Eval(%q): %v", in, err)
		}
		if actual != out {
			t.Errorf("Expected %q, Actual %q", out, actual)
		}
	}
}

func TestEvalIncr(t *testing.T) {
	cases := map[*structs.RedisCmd]string{
		&structs.RedisCmd{Cmd: "INCR", Args: nil}:            "invalid number of arguments for INCR/DECR",
		&structs.RedisCmd{Cmd: "INCR", Args: []string{"K1"}}: "OK",
	}
	for in, out := range cases {
		actual, err := Eval(in)
		if err != nil {
			t.Fatalf("Eval(%q): %v", in, err)
		}
		if actual != out {
			t.Errorf("Expected %q, Actual %q", out, actual)
		}
	}
}

func TestEvalDecr(t *testing.T) {
	cases := map[*structs.RedisCmd]string{
		&structs.RedisCmd{Cmd: "DECR", Args: nil}:            "invalid number of arguments for INCR/DECR",
		&structs.RedisCmd{Cmd: "DECR", Args: []string{"K1"}}: "OK",
	}
	for in, out := range cases {
		actual, err := Eval(in)
		if err != nil {
			t.Fatalf("Eval(%q): %v", in, err)
		}
		if actual != out {
			t.Errorf("Expected %q, Actual %q", out, actual)
		}
	}
}

func TestEvalDel(t *testing.T) {
	cases := map[*structs.RedisCmd]string{
		&structs.RedisCmd{Cmd: "DEL", Args: nil}:                                      "invalid number of keys to delete for DEL",
		&structs.RedisCmd{Cmd: "DEL", Args: []string{}}:                               "invalid number of keys to delete for DEL",
		&structs.RedisCmd{Cmd: "DEL", Args: []string{"K1", "K2"}}:                     "2",
		&structs.RedisCmd{Cmd: "DEL", Args: []string{"K1", "K2", "NOT_EXISTS", "K3"}}: "3",
		&structs.RedisCmd{Cmd: "DEL", Args: []string{"NOT_EXISTS"}}:                   "0",
	}
	for in, out := range cases {
		actual, err := Eval(in)
		if err != nil {
			t.Fatalf("Eval(%q): %v", in, err)
		}
		if actual != out {
			t.Errorf("Expected %q, Actual %q", out, actual)
		}
	}
}

func TestEvalTtl(t *testing.T) {
	cases := map[*structs.RedisCmd]string{
		&structs.RedisCmd{Cmd: "TTL", Args: nil}:                    "invalid number of arguments for TTL",
		&structs.RedisCmd{Cmd: "TTL", Args: []string{"K1"}}:         "-1",
		&structs.RedisCmd{Cmd: "TTL", Args: []string{"K2"}}:         "",
		&structs.RedisCmd{Cmd: "TTL", Args: []string{"NOT_EXISTS"}}: "-2",
	}
	for in, out := range cases {
		actual, err := Eval(in)
		if err != nil {
			t.Fatalf("Eval(%q): %v", in, err)
		}
		if actual != out {
			t.Errorf("Expected %q, Actual %q", out, actual)
		}
	}
}

func TestEvalExpire(t *testing.T) {
	cases := map[*structs.RedisCmd]string{
		&structs.RedisCmd{Cmd: "EXPIRE", Args: nil}:                           "invalid number of arguments for EXPIRE",
		&structs.RedisCmd{Cmd: "EXPIRE", Args: []string{"K1"}}:                "invalid number of arguments for EXPIRE",
		&structs.RedisCmd{Cmd: "EXPIRE", Args: []string{"K1", "-2"}}:          "invalid or out of range expiry value",
		&structs.RedisCmd{Cmd: "EXPIRE", Args: []string{"K1", "10000000000"}}: "invalid or out of range expiry value",
		&structs.RedisCmd{Cmd: "EXPIRE", Args: []string{"K1", "10"}}:          "1",
		&structs.RedisCmd{Cmd: "EXPIRE", Args: []string{"NOT_EXISTS", "10"}}:  "0",
	}
	for in, out := range cases {
		actual, err := Eval(in)
		if err != nil {
			t.Fatalf("Eval(%q): %v", in, err)
		}
		if actual != out {
			t.Errorf("Expected %q, Actual %q", out, actual)
		}
	}
}
