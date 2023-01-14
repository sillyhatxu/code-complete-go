package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	abs := func(i int) int {
		if i < 0 {
			return i * -1
		}
		return i
	}
	sort.Ints(nums)
	res, near := 0, math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			currentNear := abs(target - sum)
			if currentNear < near {
				res, near = sum, currentNear
			}
			if sum == target {
				return sum
			} else if sum < target {
				l++
			} else if sum > target {
				r--
			}
		}
	}
	return res
}

func threeSumClosest1(nums []int, target int) int {
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
