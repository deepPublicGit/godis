package structs

type RedisCmd struct {
	Cmd  string
	Args []string
}

type RedisCmds []*RedisCmd
