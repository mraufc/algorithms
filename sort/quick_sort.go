package sort

// QuickSort is a quick sort implementation that sorts a slice if ints
func QuickSort(inp []int) {
	quickSort(inp, 0, len(inp)-1)
}

func quickSort(inp []int, b, e int) {
	if e <= b {
		return
	}
	q := partition(inp, b, e)
	quickSort(inp, b, q-1)
	quickSort(inp, q+1, e)
}

func partition(inp []int, b, e int) int {
	x := inp[e]
	i := b - 1
	var t int
	for j := b; j < e; j++ {
		if inp[j] <= x {
			i++
			t = inp[i]
			inp[i] = inp[j]
			inp[j] = t
		}
	}
	i++
	t = inp[i]
	inp[i] = inp[e]
	inp[e] = t

	return i
}
