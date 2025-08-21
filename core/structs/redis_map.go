package structs

import (
	"godis/config"
	"log"
	"time"
)

var RedisStore map[string]*RedisObject

// Evictor is a callback set by the evict package to perform eviction when needed.
// It is optional to avoid cyclic imports.
var Evictor func()

func NewRedisStore() {
	RedisStore = make(map[string]*RedisObject)
}

func Get(key string) (*RedisObject, bool) {
	redisObj, ok := RedisStore[key]
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
	if len(RedisStore) >= config.MaxMemory {
		if Evictor != nil {
			Evictor()
		}
	}
	RedisStore[key] = value
}

// Del Invalid keys might be given as inputs.
func Del(key string) bool {
	if _, ok := RedisStore[key]; ok {
		delete(RedisStore, key)
		return true
	}
	return false
}

// DelExpiredKeys delete sample percentage of fixed keys, repeat until expired keys less than sample.
func DelExpiredKeys(samplePercentage int) {
	sample := float32(samplePercentage / 100)
	for deleteSample() > sample {

	}
	log.Printf("DelExpiredKeys finished: %d keys remaining", len(RedisStore))
}

// Redis calculates % of sampled deleted from a fixed limit (20)
func deleteSample() float32 {
	keysDeleted := 0
	for key, redisObject := range RedisStore {
		if redisObject.ExpiresAt != -1 && redisObject.ExpiresAt <= time.Now().UnixMilli() {
			delete(RedisStore, key)
			keysDeleted++
		}
	}
	return float32(keysDeleted) / float32(config.ExpiryLimit)
}
