package main

func removeDuplicates(nums []int) int {
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[index] {
			nums[index+1] = nums[i]
			index++
		}
	}
	return index + 1
}

func removeDuplicates1(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

func removeDuplicates2(nums []int) int {
	cur := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[cur] {
			nums[cur+1] = nums[i]
			cur++
		}
	}
	return cur + 1
}
