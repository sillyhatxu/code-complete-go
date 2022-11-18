package _62_Find_Peak_Element

//Binary Search
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		}
		return 1
	}
	l, r := 0, len(nums)-1
	for l < r {
		if l == 0 && nums[l] > nums[l+1] {
			return l
		}
		if r == len(nums)-1 && nums[r] > nums[r-1] {
			return r
		}
		if nums[l+1] > nums[l+2] && nums[l+1] > nums[l] {
			return l + 1
		}
		if nums[r-1] > nums[r-2] && nums[r-1] > nums[r] {
			return r - 1
		}
		l++
		r--
	}
	return -1
}
