package loki

import (
	"time"
)

func CurrentTimeInNs() int64 {
	return time.Now().UnixNano()
}
