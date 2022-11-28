package _79_Perfect_Squares

import "math"

//TODO
//time complexity: O(n^1/2)
func numSquares(n int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([]int, n+1, n+1)
	dp[0] = 0
	for target := 1; target < n+1; target++ {
		dp[target] = math.MaxInt32
		for i := 1; i < target+1; i++ {
			square := i * i
			if target-square < 0 {
				break
			}
			dp[target] = min(dp[target], 1+dp[target-square])
		}
	}
	return dp[n]
}
