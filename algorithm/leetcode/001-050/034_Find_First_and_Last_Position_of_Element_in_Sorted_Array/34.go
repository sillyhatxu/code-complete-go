package _34_Find_First_and_Last_Position_of_Element_in_Sorted_Array

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	findLeft := func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		for left+1 < right {
			mid := left + (right-left)/2
			if nums[mid] >= target {
				right = mid
			} else {
				left = mid
			}
		}
		if nums[left] == target {
			return left
		} else if nums[right] == target {
			return right
		} else {
			return -1
		}
	}
	findRight := func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		for left+1 < right {
			mid := left + (right-left)/2
			if nums[mid] <= target {
				left = mid
			} else {
				right = mid
			}
		}
		if nums[right] == target {
			return right
		} else if nums[left] == target {
			return left
		} else {
			return -1
		}
	}
	l := findLeft(nums, target)
	if l == -1 {
		return []int{-1, -1}
	}
	r := findRight(nums, target)
	return []int{l, r}
}

func searchRange1(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			result[0] = mid
			result[1] = mid
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if result[0] == -1 && result[1] == -1 {
		return result
	}
	for {
		if result[0] > 0 && nums[result[0]-1] == target {
			result[0] = result[0] - 1
		}
		if result[1] < len(nums)-1 && nums[result[1]+1] == target {
			result[1] = result[1] + 1
		}
		if (result[0] == 0 || nums[result[0]-1] != target) && (result[1] >= len(nums)-1 || nums[result[1]+1] != target) {
			break
		}
	}
	return result
}
