package dynamicprogramming

// SolveMatrixMultiplicationRecursive solves matrix multiplication problem with a naive recursive approach
// Three matrices ABC can be multiplied as A(BC) or (AB)C
// Those three matrices can be denoted as a slice of 4 parameters x, y, z, t
// A: x by y, B: y by z, C: z by t
// What we need to find is minimum of
// xy  + y z t + x * y * t or x y z + z t + x * z * t
// p := []int{x, y, z ,t}
// solve(p[:2]) + solve(p[1:]) + p[0] * p[1] * p[3]
// solve(p[:3]) + solve(p[2:]) + p[0] * p[2] * p[3]
func SolveMatrixMultiplicationRecursive(input []int) int {
	if len(input) < 3 {
		return 0
	}
	if len(input) == 3 {
		return input[0] * input[1] * input[2]
	}
	var min, temp int
	for i := 0; i < len(input)-2; i++ {

		temp = SolveMatrixMultiplicationRecursive(input[:i+2]) + // first part
			SolveMatrixMultiplicationRecursive(input[i+1:]) + // second part
			input[0]*input[i+1]*input[len(input)-1] // resulting matrix from first part and second part

		if i == 0 {
			min = temp
		} else {
			if temp < min {
				min = temp
			}
		}
	}
	return min
}

// SolveMatrixMultiplicationDP solves the problem by using a slice of slices for storing previously calculated values.
func SolveMatrixMultiplicationDP(input []int) int {
	m := make([][]int, len(input)+1)
	for i := range m {
		m[i] = make([]int, len(input)+1)
	}
	return solveMatrixMultiplicationDPAux(input, 0, len(input), m)
}

func solveMatrixMultiplicationDPAux(input []int, k, l int, m [][]int) int {
	if m[k][l] > 0 {
		return m[k][l]
	}
	if l-k < 3 {
		return 0
	}
	if l-k == 3 {
		res := input[k] * input[k+1] * input[k+2]
		m[k][l] = res
		return res
	}

	var min, temp, first, second int

	for i := k; i < l-2; i++ {
		first = solveMatrixMultiplicationDPAux(input, k, i+2, m)
		second = solveMatrixMultiplicationDPAux(input, i+1, l, m)
		temp = first + second + input[k]*input[i+1]*input[l-1]
		if i == k {
			min = temp
		} else {
			if temp < min {
				min = temp
			}
		}
	}

	m[k][l] = min

	return min
}
