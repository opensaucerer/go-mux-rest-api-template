package helper

import (
	"time"
)

func GetTimeInFuture(duration int) time.Time {
	return time.Now().Add(time.Duration(duration) * time.Second)
}

func GetTimeInPast(duration int) time.Time {
	return time.Now().Add(time.Duration(-duration) * time.Second)
}
