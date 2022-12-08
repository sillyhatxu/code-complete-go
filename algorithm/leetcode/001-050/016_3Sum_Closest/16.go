package main

import (
	"math"
	"sort"
)

func threeSumClosest1(nums []int, target int) int {
	abs := func(i int) int {
		if i < 0 {
			return i * -1
		}
		return i
	}
	sort.Ints(nums)

	n, SumClosest := len(nums), nums[0]+nums[1]+nums[2] // 初始化为前三元素的值，避免任何个

	for i := 0; i < n-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}
			if abs(sum-target) < abs(SumClosest-target) {
				SumClosest = sum
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}

	return SumClosest
}
func threeSumClosest(nums []int, target int) int {
	abs := func(i int) int {
		if i < 0 {
			return i * -1
		}
		return i
	}
	sort.Ints(nums)
	res, near := 0, math.MaxInt64 //initial near
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			tempNear := abs(target - sum)
			if tempNear < near {
				near = tempNear
				res = sum
			}
			if target == sum {
				return sum
			} else if target > sum {
				left++
			} else { //target < sum
				right--
			}
		}

	}
	return res
}
