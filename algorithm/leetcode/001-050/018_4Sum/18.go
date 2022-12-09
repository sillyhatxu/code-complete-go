package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int
	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, length-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left+1 < right && nums[left] == nums[left+1] {
						left++
					}
					left++
					for left < right-1 && nums[right] == nums[right-1] {
						right--
					}
					right--
				} else if sum > target {
					right--
				} else {
					//sum < target
					left++
				}
			}
		}
	}
	return result
}
