package climbingStairs

func ClimbStairs(n int) int {
	if n <= 1 {
		return n
	}

	// Create an array to store the number of ways to reach each step
	ways := make([]int, n+1)

	// Base cases
	ways[0] = 1 // 1 way to stay at ground level (do nothing)
	ways[1] = 1 // 1 way to reach the first step (1 step)

	// Fill the ways array
	for i := 2; i <= n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}

	return ways[n]

}
