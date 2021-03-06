package sort

import (
	"algorithms/utils"
	"testing"
)

func TestCountingSort(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{
			name: "myList",
			data: []int{31, 41, 59, 26, 41, 58},
		},
		{
			name: "myList2",
			data: []int{31, 41, 59, -26, 41, -58},
		},
		{
			name: "5 random integers",
			data: utils.RandomInts(5),
		},
		{
			name: "10 random integers",
			data: utils.RandomInts(10),
		},
		{
			name: "50 random integers",
			data: utils.RandomInts(50),
		},
		{
			name: "100 random integers",
			data: utils.RandomInts(100),
		},
		{
			name: "750 random integers",
			data: utils.RandomInts(750),
		},
		{
			name: "1000 random integers",
			data: utils.RandomInts(1000),
		},
		{
			name: "10000 random integers",
			data: utils.RandomInts(10000),
		},
		{
			name: "100000 random integers",
			data: utils.RandomInts(100000),
		},
		{
			name: "120000 random integers",
			data: utils.RandomInts(120000),
		},
		{
			name: "140000 random integers",
			data: utils.RandomInts(140000),
		},
		{
			name: "160000 random integers",
			data: utils.RandomInts(160000),
		},
		{
			name: "180000 random integers",
			data: utils.RandomInts(180000),
		},
		{
			name: "200000 random integers",
			data: utils.RandomInts(200000),
		},
		{
			name: "2000000 random integers",
			data: utils.RandomInts(2000000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.data = CountingSort(tt.data)
			for i := 1; i < len(tt.data); i++ {
				if tt.data[i] < tt.data[i-1] {
					t.Log("test case", tt.name, "failed", tt.data[:i+1], "reason:", tt.data[i], "is less than", tt.data[i-1])
					t.FailNow()
				}
			}
		})
	}
}
