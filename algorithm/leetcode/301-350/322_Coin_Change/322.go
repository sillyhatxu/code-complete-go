package _22_Coin_Change

import "fmt"

func coinChange(coins []int, amount int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	total := amount + 1
	dp := make([]int, total, total)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = total
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}
	if dp[amount] == total {
		return -1
	}
	return dp[amount]
}

func coinChange1(coins []int, amount int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	total := amount + 1
	dp := make([]int, total, total)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = total
		for _, coin := range coins {
			if (i-coin) >= 0 && dp[i-coin] != total {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
		fmt.Println(dp)
	}
	if dp[amount] == total {
		return -1
	}
	return dp[amount]
}
