package sort

// CountingSort is a counting sort implementation that sorts slice of ints
func CountingSort(input []int) []int {
	min, max := findMinAndMax(input)
	temp := make([]int, max-min+1)
	result := make([]int, len(input))

	for _, v := range input {
		temp[v-min]++
	}
	for ix := 1; ix < len(temp); ix++ {
		temp[ix] += temp[ix-1]
	}
	for ij := len(input) - 1; ij >= 0; ij-- {
		result[temp[input[ij]-min]-1] = input[ij]
		temp[input[ij]-min]--
	}
	return result
}

func findMinAndMax(input []int) (min, max int) {
	if len(input) == 0 {
		panic("invalid input slice length 0")
	}
	max, min = input[0], input[0]
	for _, v := range input {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}
