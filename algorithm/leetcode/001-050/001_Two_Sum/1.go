package main

import "fmt"

func twoSum(nums []int, target int) []int {
	targetMap := make(map[int]int, len(nums))
	for i := range nums {
		if index, ok := targetMap[nums[i]]; ok {
			return []int{index, i}
		} else {
			targetMap[target-nums[i]] = i
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 18))
	fmt.Println(twoSum(nums, 18))
}
