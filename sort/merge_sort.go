package sort

func merge(left, right []int) []int {
	x := make([]int, len(left)+len(right))
	il, ir, ix := 0, 0, 0
	if left[il] < right[ir] {
		x[ix] = left[il]
		il++
	} else {
		x[ix] = right[ir]
		ir++
	}
	ix++
	for ; ix < len(x); ix++ {
		switch {
		case il >= len(left) && ir >= len(right):
			break
		case il >= len(left):
			x[ix] = right[ir]
			ir++
		case ir >= len(right):
			x[ix] = left[il]
			il++
		case left[il] >= right[ir]:
			x[ix] = right[ir]
			ir++
		case left[il] < right[ir]:
			x[ix] = left[il]
			il++
		}
	}
	return x
}

// MergeSort is a merge sort implementation that sorts a slice of ints
func MergeSort(inp []int) []int {
	if len(inp) == 1 {
		return inp
	}
	m := len(inp) >> 1
	l := inp[:m]
	r := inp[m:]
	return merge(MergeSort(l), MergeSort(r))
}
