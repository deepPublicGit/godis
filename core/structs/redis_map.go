package structs

import (
	"godis/config"
	"log"
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

// DelExpiredKeys delete sample percentage of fixed keys, repeat until expired keys less than sample.
func DelExpiredKeys(samplePercentage int) {
	sample := float32(samplePercentage / 100)
	for deleteSample() > sample {

	}
	log.Printf("DelExpiredKeys finished: %d keys remaining", len(redisMap))
}

// Redis calculates % of sampled deleted from a fixed limit (20)
func deleteSample() float32 {
	keysDeleted := 0
	for key, redisObject := range redisMap {
		if redisObject.ExpiresAt != -1 && redisObject.ExpiresAt <= time.Now().UnixMilli() {
			delete(redisMap, key)
			keysDeleted++
		}
	}
	return float32(keysDeleted) / float32(config.ExpiryLimit)
}
