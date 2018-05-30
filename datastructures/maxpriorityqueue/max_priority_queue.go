package maxpriorityqueue

import (
	"fmt"
)

type heap struct {
	data []int
	size int
}

// MaxPriorityQueue is a max heap based max priority queue impelemntation
type MaxPriorityQueue heap

func (q *MaxPriorityQueue) String() string {
	return fmt.Sprint(q.data[:q.size])
}

// basic operations are modified for 0 indexed heap
func parent(n int) int {
	return ((n + 1) >> 1) - 1
}
func left(n int) int {
	return ((n + 1) << 1) - 1
}
func right(n int) int {
	return left(n) + 1
}

// NewMaxPriorityQueue returns a new maximum priority queue
func NewMaxPriorityQueue(input []int) *MaxPriorityQueue {
	q := &MaxPriorityQueue{data: input, size: len(input)}
	for i := (len(input) >> 1) - 1; i >= 0; i-- {
		q.maxHeapify(i)
	}
	return q
}

// Insert inserts a new element to the queue
func (q *MaxPriorityQueue) Insert(k int) {
	q.size++
	q.data = append(q.data, k-1)
	q.IncreaseKey(q.size-1, k)
}

// Maximum returns the maximum element, i.e. the first element in the queue
func (q *MaxPriorityQueue) Maximum() int {
	if q.size < 1 {
		panic("heap underflow")
	}
	return q.data[0]
}

// ExtractMaximum extracts and returns the maximum element, and 'heapifies' the remaining elements
func (q *MaxPriorityQueue) ExtractMaximum() int {
	max := q.Maximum()
	q.data[0] = q.data[q.size-1]
	q.size--
	q.maxHeapify(0)
	return max
}

// IncreaseKey increases the item at index position i to k only if k is greater than the value of the item.
func (q *MaxPriorityQueue) IncreaseKey(i, k int) {
	if q.data[i] >= k {
		return
	}
	q.data[i] = k
	var temp int
	for i > 0 && q.data[parent(i)] < q.data[i] {
		temp = q.data[i]
		q.data[i] = q.data[parent(i)]
		q.data[parent(i)] = temp
		i = parent(i)
	}
}

func (q *MaxPriorityQueue) maxHeapify(i int) {
	l := left(i)
	r := right(i)
	largest := i
	if l < q.size && q.data[l] > q.data[i] {
		largest = l
	}
	if r < q.size && q.data[r] > q.data[largest] {
		largest = r
	}
	if largest == i {
		return
	}
	t := q.data[i]
	q.data[i] = q.data[largest]
	q.data[largest] = t
	q.maxHeapify(largest)
}
