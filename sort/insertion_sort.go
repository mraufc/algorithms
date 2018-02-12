package sort

// InsertionSort is an insertion sort implementation that sorts a slice of ints
func InsertionSort(data []int) {
	var key, i int
	for j := 1; j < len(data); j++ {
		key = data[j]
		i = j - 1
		for i >= 0 && data[i] > key {
			data[i+1] = data[i]
			i--
		}
		data[i+1] = key
	}
}