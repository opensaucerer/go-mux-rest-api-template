package helper

import (
	"math/rand"
)

func PickAtRandom(arr []string) string {
	return arr[rand.Intn(len(arr))]
}
