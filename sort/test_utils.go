package sort

import (
	"math/rand"
	"time"
)

func randomInts(size int) []int {
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