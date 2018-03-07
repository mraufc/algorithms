package dataStructures

type heap struct {
	data []int
	size int
}

type maxPriorityQueue struct {
	data []int
	size int
}

// operations
func parent(n int) int {
	return n >> 1
}
func left(n int) int {
	return n << 1
}
func right(n int) int {
	return (n << 1) + 1
}

func buildMaxPriorityQueue(input []int) *maxPriorityQueue {
	q := &maxPriorityQueue{data: input, size: len(input)}
	for i := (len(input) >> 1) - 1; i >= 0; i-- {
		q.maxHeapify(i)
	}
	return q
}

func (q *maxPriorityQueue) insert(k int) {
	// TODO
	q.size++
	q.data = append(q.data, k-1)
	q.increaseKey(q.size-1, k)
}

func (q *maxPriorityQueue) maximum() int {
	if q.size < 1 {
		panic("heap underflow")
	}
	return q.data[0]
}

func (q *maxPriorityQueue) extractMaximum() int {
	max := q.maximum()
	q.data[0] = q.data[q.size-1]
	q.size--
	q.maxHeapify(0)
	return max
}

func (q *maxPriorityQueue) increaseKey(i, k int) {
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

func (q *maxPriorityQueue) maxHeapify(i int) {
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
