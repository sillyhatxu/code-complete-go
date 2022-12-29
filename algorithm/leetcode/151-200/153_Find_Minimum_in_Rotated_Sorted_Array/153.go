package _53_Find_Minimum_in_Rotated_Sorted_Array

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l+1 < r {
		mid := (r-l)/2 + l
		if nums[mid] > nums[r] {
			l = mid
		} else if nums[mid] < nums[r] {
			r = mid
		}
	}
	if nums[l] < nums[r] {
		return nums[l]
	}
	return nums[r]
}
