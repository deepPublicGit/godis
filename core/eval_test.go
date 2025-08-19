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
