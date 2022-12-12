package _33_Search_in_Rotated_Sorted_Array

func search(nums []int, target int) int {
	if len(nums) == 0 || (len(nums) == 1 && nums[0] != target) {
		return -1
	} else if len(nums) == 1 && nums[0] == target {
		return 0
	}
	pivot := findPivotIndex(nums)
	if target >= nums[0] && target <= nums[pivot] {
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
		//[0,1,2,3,4,5]
		return len(nums) - 1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			//eg: 6,7,8,9 | 3,4,5  mid = 3 ->  nums[3] < nums[3+1] -> 9 < 3 return mid(3)
			return mid
		} else if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}
