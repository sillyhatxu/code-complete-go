package _00_Longest_Increasing_Subsequence

import "math"

func lengthOfLIS(nums []int) int {
	tails := make([]int, len(nums)+1, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		tails[i] = math.MaxInt64
	}
	for i := 0; i < len(nums); i++ {
		index := helper(tails, nums[i])
		tails[index] = nums[i]
	}
	res := 0
	for i := 0; i < len(tails); i++ {
		if tails[i] != math.MaxInt64 {
			res++
		}
	}
	return res
}
func helper(tails []int, target int) int {
	start, end := 0, len(tails)-1
	for start+1 < end {
		mid := (start + end) / 2
		if tails[mid] == target {
			return mid
		} else if tails[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}
	return end
}
