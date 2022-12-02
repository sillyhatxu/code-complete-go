package _34_Increasing_Triplet_Subsequence

import "math"

func increasingTriplet(nums []int) bool {
	min, mid := math.MaxInt32, math.MaxInt32
	for _, current := range nums {
		if current <= min {
			//case 1
			min = current
		} else if current <= mid {
			//case 2
			mid = current
		} else if current > mid {
			//case 3
			return true
		}
	}
	return false
}
