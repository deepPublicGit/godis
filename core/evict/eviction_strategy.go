package evict

import (
	"godis/config"
	"godis/core/structs"
)

func init() {
	// Register eviction callback to avoid cyclic imports.
	structs.Evictor = Evict
}

func Evict() {
	switch config.EvictionStrategy {
	case "random":
		randomEvict()
	default:
		// no-op for now
	}
}
