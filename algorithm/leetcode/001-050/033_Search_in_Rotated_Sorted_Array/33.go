package _33_Search_in_Rotated_Sorted_Array

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 && nums[0] == target {
		return 0
	} else if len(nums) == 1 && nums[0] != target {
		return -1
	}
	pivot := findPivotIndex(nums)
	if pivot >= 0 && target >= nums[0] && target <= nums[pivot] {
		return binarySearch(nums, 0, pivot, target)
	} else {
		return binarySearch(nums, pivot+1, len(nums)-1, target)
	}
}

func binarySearch(nums []int, left, right, target int) int {
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func findPivotIndex(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return len(nums) - 1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			return mid
		} else if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}
