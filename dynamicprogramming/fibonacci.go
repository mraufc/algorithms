package dynamicprogramming

// SolveFibonacciRecursive returns n'th (0 indexed) fibonacci number by using a simple recursive approach
func SolveFibonacciRecursive(n int64) int64 {
	if n < 0 {
		panic("invalid index")
	}
	if n <= 1 {
		return n
	}
	return SolveFibonacciRecursive(n-1) + SolveFibonacciRecursive(n-2)
}

// SolveFibonacciDP eturns n'th (0 indexed) fibonacci number by using an array to store pre-computed values
func SolveFibonacciDP(n int64) int64 {
	if n < 0 {
		panic("invalid index")
	}
	if n <= 1 {
		return n
	}
	f := make([]int64, n+1)
	f[0] = 0
	f[1] = 1
	return solveFibonacciDPAux(n-1, f) + solveFibonacciDPAux(n-2, f)
}

func solveFibonacciDPAux(n int64, f []int64) int64 {
	if n <= 1 {
		return f[n]
	}
	if n > 1 && f[n] > 0 {
		return f[n]
	}
	r := solveFibonacciDPAux(n-1, f) + solveFibonacciDPAux(n-2, f)
	f[n] = r
	return r
}

// SolveFibonacciDPSO returns n'th (0 indexed) fibonacci number by using an array to store previous two pre-computed values
// DPSO stands for Dynamic Programming & Space Optimized
func SolveFibonacciDPSO(n int64) int64 {
	if n < 0 {
		panic("invalid index")
	}
	if n <= 1 {
		return n
	}
	f := make([]int64, 2)
	f[0] = 0
	f[1] = 1
	for i := int64(2); i < n; i++ {
		f[0], f[1] = f[1], f[0]+f[1]
	}
	return f[0] + f[1]
}
