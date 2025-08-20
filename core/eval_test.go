package core

import (
	"godis/core/structs"
	"testing"
)

func TestEvalPingValidCases(t *testing.T) {
	cases := map[*structs.RedisCommands]string{
		&structs.RedisCommands{Cmd: "PING", Args: nil}:            "PONG",
		&structs.RedisCommands{Cmd: "PING", Args: []string{"OK"}}: "PONG OK",
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
	cases := map[*structs.RedisCommands]string{
		&structs.RedisCommands{Cmd: "PONG", Args: nil}:                       "invalid command",
		&structs.RedisCommands{Cmd: "PING", Args: []string{"OKIE", "DOKIE"}}: "invalid number of arguments",
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
	cases := map[*structs.RedisCommands]string{
		&structs.RedisCommands{Cmd: "GET", Args: nil}:            "invalid number of arguments",
		&structs.RedisCommands{Cmd: "GET", Args: []string{"K1"}}: "V1",
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
	cases := map[*structs.RedisCommands]string{
		&structs.RedisCommands{Cmd: "SET", Args: nil}:                  "invalid number of arguments",
		&structs.RedisCommands{Cmd: "SET", Args: []string{"K1", "10"}}: "PONG OK",
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
	cases := map[*structs.RedisCommands]string{
		&structs.RedisCommands{Cmd: "INCR", Args: nil}:            "PONG",
		&structs.RedisCommands{Cmd: "INCR", Args: []string{"K1"}}: "PONG OK",
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
	cases := map[*structs.RedisCommands]string{
		&structs.RedisCommands{Cmd: "DECR", Args: nil}:            "PONG",
		&structs.RedisCommands{Cmd: "DECR", Args: []string{"K1"}}: "PONG OK",
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
	cases := map[*structs.RedisCommands]string{
		&structs.RedisCommands{Cmd: "DEL", Args: nil}:                                      "invalid number of keys to delete for DEL",
		&structs.RedisCommands{Cmd: "DEL", Args: []string{}}:                               "invalid number of keys to delete for DEL",
		&structs.RedisCommands{Cmd: "DEL", Args: []string{"K1", "K2"}}:                     "2",
		&structs.RedisCommands{Cmd: "DEL", Args: []string{"K1", "K2", "NOT_EXISTS", "K3"}}: "3",
		&structs.RedisCommands{Cmd: "DEL", Args: []string{"NOT_EXISTS"}}:                   "0",
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
	cases := map[*structs.RedisCommands]string{
		&structs.RedisCommands{Cmd: "TTL", Args: nil}:            "invalid number of arguments for TTL",
		&structs.RedisCommands{Cmd: "TTL", Args: []string{"K1"}}: "PONG OK",
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
	cases := map[*structs.RedisCommands]string{
		&structs.RedisCommands{Cmd: "EXPIRE", Args: nil}:                           "invalid number of arguments for EXPIRE",
		&structs.RedisCommands{Cmd: "EXPIRE", Args: []string{"K1"}}:                "invalid number of arguments for EXPIRE",
		&structs.RedisCommands{Cmd: "EXPIRE", Args: []string{"K1", "-2"}}:          "invalid or out of range expiry value",
		&structs.RedisCommands{Cmd: "EXPIRE", Args: []string{"K1", "10000000000"}}: "invalid or out of range expiry value",
		&structs.RedisCommands{Cmd: "EXPIRE", Args: []string{"K1", "10"}}:          "1",
		&structs.RedisCommands{Cmd: "EXPIRE", Args: []string{"NOT_EXISTS", "10"}}:  "0",
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
