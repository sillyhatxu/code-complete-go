package main

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			//nums[mid] < target
			l = mid + 1
		}
	}
	return l
}

//O(n)
func searchInsertON(nums []int, target int) int {
	previous := 0
	for i, v := range nums {
		if v >= target {
			return i
		}
		previous++
	}
	return previous
}
