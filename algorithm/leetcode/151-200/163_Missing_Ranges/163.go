package _63_Missing_Ranges

import "fmt"

func findMissingRanges(nums []int, lower int, upper int) []string {
	if len(nums) == 0 {
		if lower == upper {
			return []string{fmt.Sprintf("%d", lower)}
		} else {
			return []string{fmt.Sprintf("%d->%d", lower, upper)}
		}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	lower = min(nums[0], lower)
	var res []string
	for i := 0; i < len(nums); i++ {
		if lower == nums[i] {
			lower++
		} else if lower < nums[i] {
			if lower == nums[i]-1 {
				res = append(res, fmt.Sprintf("%d", lower))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", lower, nums[i]-1))
			}
			lower = nums[i] + 1
		}
	}
	if upper > nums[len(nums)-1] {
		if nums[len(nums)-1]+1 == upper {
			res = append(res, fmt.Sprintf("%d", upper))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[len(nums)-1]+1, upper))
		}
	}
	return res
}
