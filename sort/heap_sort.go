package sort

type heap struct {
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

func maxHeapify(h *heap, i int) {
	l := left(i)
	r := right(i)
	largest := i
	if l < h.size && h.data[l] > h.data[i] {
		largest = l
	}
	if r < h.size && h.data[r] > h.data[largest] {
		largest = r
	}
	if largest == i {
		return
	}
	t := h.data[i]
	h.data[i] = h.data[largest]
	h.data[largest] = t
	maxHeapify(h, largest)
}

func buildMaxHeap(input []int) *heap {
	h := &heap{data: input, size: len(input)}
	for i := (len(input) >> 1) - 1; i >= 0; i-- {
		maxHeapify(h, i)
	}
	return h
}

// HeapSort is a heap sort implementation that sorts a slice of ints
func HeapSort(input []int) {
	h := buildMaxHeap(input)
	var t int
	for i := len(h.data) - 1; i >= 1; i-- {
		t = h.data[0]
		h.data[0] = h.data[i]
		h.data[i] = t
		h.size--
		maxHeapify(h, 0)
	}
}
