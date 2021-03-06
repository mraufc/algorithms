package utils

import (
	"math/rand"
	"time"
)

func RandomInts(size int) []int {
	rand.Seed(time.Now().Unix())
	if size <= 0 {
		size = 0
	}
	r := make([]int, size)
	for ix := range r {
		r[ix] = rand.Intn(size)
	}
	return r
}
