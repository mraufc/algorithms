package datastructures

import (
	"testing"
)

func TestMaxPriorityQueue(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{
			"List_1",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			"List_2",
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			"List_3",
			[]int{9, 7, 5, 3, 1, 2, 4, 6, 8, 10},
		},
		{
			"List_4",
			[]int{1, 3, 5, 7, 9, 10, 8, 6, 4, 2},
		},
		{
			"List_5",
			[]int{10, 1, 9, 2, 8, 3, 7, 4, 6, 5},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := NewMaxPriorityQueue(tc.data)
			t.Log(h)
			for i := 11; i <= 13; i++ {
				h.Insert(i)
				t.Log(h)
			}
			for i := 13; i >= 1; i-- {
				t.Log(h)
				if i != h.ExtractMaximum() {
					t.FailNow()
				}
			}
		})
	}
}
