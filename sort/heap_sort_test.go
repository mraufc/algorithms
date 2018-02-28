package sort

import "testing"

func TestHeapSort(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{
			name: "myList",
			data: []int{31, 41, 59, 26, 41, 58},
		},
		{
			name: "5 random integers",
			data: randomInts(5),
		},
		{
			name: "10 random integers",
			data: randomInts(10),
		},
		{
			name: "50 random integers",
			data: randomInts(50),
		},
		{
			name: "100 random integers",
			data: randomInts(100),
		},
		{
			name: "750 random integers",
			data: randomInts(750),
		},
		{
			name: "1000 random integers",
			data: randomInts(1000),
		},
		{
			name: "10000 random integers",
			data: randomInts(10000),
		},
		{
			name: "100000 random integers",
			data: randomInts(100000),
		},
		{
			name: "120000 random integers",
			data: randomInts(120000),
		},
		{
			name: "140000 random integers",
			data: randomInts(140000),
		},
		{
			name: "160000 random integers",
			data: randomInts(160000),
		},
		{
			name: "180000 random integers",
			data: randomInts(180000),
		},
		{
			name: "200000 random integers",
			data: randomInts(200000),
		},
		{
			name: "2000000 random integers",
			data: randomInts(2000000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.data)
			for i := 1; i < len(tt.data); i++ {
				if tt.data[i] < tt.data[i-1] {
					t.Log("test case", tt.name, "failed", tt.data[:i+1], "reason:", tt.data[i], "is less than", tt.data[i-1])
					t.FailNow()
				}
			}
		})
	}
}
