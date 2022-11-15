package helper

import (
	"time"
)

func GetCurrentTime() time.Time {
	return time.Now()
}

func GetCurrentTimeAsUnix() int64 {
	return time.Now().Unix()
}
