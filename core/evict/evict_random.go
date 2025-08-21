package evict

import (
	"godis/config"
	"godis/core/structs"
	"log"
)

func randomEvict() {
	keysDeleted := 0
	for key := range structs.RedisStore {
		delete(structs.RedisStore, key)
		keysDeleted++
		if len(structs.RedisStore) < config.MaxMemory {
			log.Printf("Evicted %d key(s)", keysDeleted)
			return
		}
	}
}
