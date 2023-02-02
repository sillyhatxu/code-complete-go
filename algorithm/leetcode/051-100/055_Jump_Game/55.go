package _5_Jump_Game

func canJump(nums []int) bool {
	right := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < len(nums); i++ {
		if i > right {
			return false
		}
		right = max(right, i+nums[i])
	}
	return right >= len(nums)-1
}

func canJump1(nums []int) bool {
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
