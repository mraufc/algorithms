package dynamicprogramming

// problem: you're given a slice of integers and you can choose to place either an addition (+) or multiplication (*) operator
// between each number. also you can place parentheses anywhere you like. find the maximum value you can achieve.

// SolveMaxNumberWithSignsRecursive solves the above problem with a naive recursive approach.
func SolveMaxNumberWithSignsRecursive(input []int) int {
	if len(input) == 0 {
		return 0
	}
	if len(input) == 1 {
		return input[0]
	}
	if len(input) == 2 {
		m := input[0] * input[1]
		s := input[0] + input[1]
		if m > s {
			return m
		}
		return s
	}
	max := 0
	for i := 1; i < len(input); i++ {
		l := SolveMaxNumberWithSignsRecursive(input[:i])
		r := SolveMaxNumberWithSignsRecursive(input[i:])
		res := SolveMaxNumberWithSignsRecursive([]int{l, r})
		if res > max {
			max = res
		}
	}
	return max
}

// SolveMaxNumberWithSignsDP uses a slice of slices to remember the solutions for sub-slices to solve the problem
func SolveMaxNumberWithSignsDP(input []int) int {
	m := make([][]int, len(input)+1)
	for i := range m {
		m[i] = make([]int, len(input)+1)
	}
	return solveMaxNumberWithSignsDPAux(input, 0, len(input), m)
}

func solveMaxNumberWithSignsDPAux(input []int, k, l int, m [][]int) int {
	if l-k == 0 {
		return 0
	}
	if m[k][l] > 0 {
		return m[k][l]
	}
	if l-k == 1 {
		m[k][l] = input[k]
		return input[k]
	}
	if l-k == 2 {
		mul := input[k] * input[k+1]
		add := input[k] + input[k+1]
		if mul > add {
			m[k][l] = mul
			return mul
		}
		m[k][l] = add
		return add
	}
	max := 0
	for i := k + 1; i < l; i++ {
		left := solveMaxNumberWithSignsDPAux(input, k, i, m)
		right := solveMaxNumberWithSignsDPAux(input, i, l, m)
		mul := left * right
		add := left + right
		if mul > max {
			max = mul
		}
		if add > max {
			max = add
		}
	}
	m[k][l] = max
	return max
}
