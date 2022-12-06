package _45_Jump_Game_II

func jump(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	jumpTime, end, farthest := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(i+nums[i], farthest)
		if farthest >= len(nums)-1 {
			jumpTime++
			break
		}
		if i == end {
			jumpTime++
			end = farthest
		}
	}
	return jumpTime
}

func jumpDP(nums []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = i
	}
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j >= len(nums) {
				break
			}
			dp[i+j] = min(dp[i+j], dp[i]+1)
		}
	}
	return dp[len(nums)-1]
}
