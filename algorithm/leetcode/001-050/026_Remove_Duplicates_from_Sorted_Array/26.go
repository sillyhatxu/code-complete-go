package main

func removeDuplicates1(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

func removeDuplicates(nums []int) int {
	cur := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[cur] {
			nums[cur+1] = nums[i]
			cur++
		}
	}
	return cur + 1
}
