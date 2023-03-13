package ddd

func test(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxCurrent := nums[0]
	maxEnd := nums[0]
	for i := 1; i < len(nums); i++ {
		maxEnd = max(nums[i], maxEnd+nums[i])
		maxCurrent = max(maxCurrent, maxEnd)
	}
	return maxCurrent
}
