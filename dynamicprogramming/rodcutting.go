package dynamicprogramming

// Introduction to Algorithms 3rd Edition

// Problem: Serling Enterprises buys long steel rods and cuts them into shorter rods,
// which it then sells. Each cut is free.
// The management of Serling Enterprises wants to know the best way to cut up the rods.

// SolveRodCuttingMemoized uses top down approach to solve rod-cutting problem given a slice of prices for each rod of length of i
// with memoization added to recursive solution.
// price of a rod of length i is stored in prices[i-1]
// length is the length of the rod
func SolveRodCuttingMemoized(prices []int, length int) int {
	if length <= 0 {
		return 0
	}
	mem := make([]int, length)
	return solveRodCuttingMemoizedAux(prices, mem, length)
}

func solveRodCuttingMemoizedAux(prices, mem []int, length int) int {
	if length <= 0 {
		return 0
	}
	if mem[length-1] > 0 {
		return mem[length-1]
	}
	var q, r int
	for i := 1; i <= length; i++ {
		r = prices[i-1] + solveRodCuttingMemoizedAux(prices, mem, length-i)
		if r > q {
			q = r
		}
	}
	mem[length-1] = q
	return q
}

// SolveRodCuttingTabulated uses bottom up approach to solve rod-cutting problem given a slice of prices for each rod of length i
// price of a rod of length i is stored in prices[i-1]
// length is the length of the rod
func SolveRodCuttingTabulated(prices []int, length int) int {
	r := make([]int, length)
	r[0] = prices[0]
	var q, v int
	for i := 1; i < length; i++ {
		q = prices[i] // default is no cuts
		for j := 1; j < i; j++ {
			v = prices[j-1] + r[i-j]
			if v > q {
				q = v
			}
		}
		r[i] = q
	}
	return r[length-1]
}

// SolveRodCuttingRecursiveTopDown solves rod-cutting problem given a slice of prices for each length of i with a recursive top down approach.
// price of a rod of length i is stored in prices[i-1]
// length is the length of the rod
func SolveRodCuttingRecursiveTopDown(prices []int, length int) int {
	if length <= 0 {
		return 0
	}
	var q, r int
	for i := 1; i <= length; i++ {
		r = prices[i-1] + SolveRodCuttingRecursiveTopDown(prices, length-i)
		if r > q {
			q = r
		}
	}
	return q
}
