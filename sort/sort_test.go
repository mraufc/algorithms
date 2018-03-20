package sort

import (
	"algorithms/utils"
	"testing"
)

func TestCompareSortAlgorithms(t *testing.T) {
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
			data: []int{31, 59, 41, 59, 26, 41, 58},
		},
		{
			name: "myList3",
			data: []int{31, 31, 41, 1, 26, 41, 59},
		},
		{
			name: "myList4",
			data: []int{1, 3, 3, 3, 4},
		},
		{
			name: "myList5",
			data: []int{5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 8, 1, 3, 3, 3, 4, 8, 8, 8, 8, 7},
		},
		{
			name: "myList6",
			data: []int{31, 59, 41, -59, 26, 41, 58, 1, 2, 3, -4, -5, -6, 7, 8, 100, -101},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSortData := make([]int, len(tt.data))
			copy(quickSortData, tt.data)
			heapSortData := make([]int, len(tt.data))
			copy(heapSortData, tt.data)
			mergeSortData := make([]int, len(tt.data))
			copy(mergeSortData, tt.data)
			countingSortData := CountingSort(tt.data)
			InsertionSort(tt.data)
			mergeSortData = MergeSort(mergeSortData)
			HeapSort(heapSortData)
			QuickSort(quickSortData)
			for i := 1; i < len(tt.data); i++ {
				if tt.data[i] < tt.data[i-1] {
					t.Log("insertion sort for test case", tt.name, "failed", tt.data[:i+1], "reason:", tt.data[i], "is less than", tt.data[i-1])
					t.FailNow()
				}
			}
			for i := 0; i < len(tt.data); i++ {
				if tt.data[i] != mergeSortData[i] || tt.data[i] != heapSortData[i] || tt.data[i] != quickSortData[i] || tt.data[i] != countingSortData[i] {
					t.Log("test case", tt.name, "failed")
					t.Log("position", i)
					s := 0
					e := len(tt.data)
					if i-5 > s {
						s = i - 5
					}
					if i+5 < e {
						e = i + 5
					}
					t.Log("insertion sort data", tt.data[s:e])
					t.Log("merge sort data", mergeSortData[s:e])
					t.Log("heap sort data", heapSortData[s:e])
					t.Log("quick sort data", quickSortData[s:e])
					t.Log("counting sort data", countingSortData[s:e])
					t.FailNow()
				}
			}
		})
	}
}
