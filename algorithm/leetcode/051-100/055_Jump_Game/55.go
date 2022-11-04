package _5_Jump_Game

func canJump(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		max = mathMax(max, i+nums[i])
	}
	return max >= -len(nums)-1
}

func mathMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
