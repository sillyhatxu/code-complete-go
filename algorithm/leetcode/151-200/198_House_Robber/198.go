package _98_House_Robber

func rob(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make(map[int]int)
	dp[0] = nums[0]
	if len(nums) == 1 {
		return dp[0]
	}
	dp[1] = max(nums[0], nums[1])
	if len(nums) == 2 {
		return dp[1]
	}
	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[len(nums)-1]
}
