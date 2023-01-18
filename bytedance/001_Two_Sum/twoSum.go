package main

func twoSum(nums []int, target int) []int {
	temp := make(map[int]int)
	for i := range nums {
		if value, ok := temp[nums[i]]; ok {
			return []int{value, i}
		}
		temp[target-nums[i]] = i
	}
	return nil
}
