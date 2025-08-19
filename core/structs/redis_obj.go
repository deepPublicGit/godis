package structs

import "time"

type RedisObject struct {
	Value     any
	ExpiresAt int64
}

func NewRedisObject(val any, expMs int64) *RedisObject {
	expiresAt := int64(-1)
	if expMs > 0 {
		expiresAt = time.Now().UnixMilli() + expMs //32 bit would cause Y2K38
	}
	return &RedisObject{
		Value:     val,
		ExpiresAt: expiresAt,
	}
}
