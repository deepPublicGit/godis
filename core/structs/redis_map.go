package structs

import (
	"time"
)

var redisMap map[string]*RedisObject

func NewRedisMap() {
	redisMap = make(map[string]*RedisObject)
}

func Get(key string) (*RedisObject, bool) {
	redisObj, ok := redisMap[key]
	if !ok {
		return nil, false
	}
	if redisObj.ExpiresAt > time.Now().UnixMilli() {
		Del(key)
		return nil, false
	}
	return redisObj, true
}

func Set(key string, value *RedisObject) {
	redisMap[key] = value
}

// Del Invalid keys might be given as inputs.
func Del(key string) bool {
	if _, ok := redisMap[key]; ok {
		delete(redisMap, key)
		return true
	}
	return false
}
