package utils

import "time"

func GetExpiryInUnixMs(expMs int64) int64 {
	return time.Now().UnixMilli() + expMs
}
