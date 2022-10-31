package _34_Find_First_and_Last_Position_of_Element_in_Sorted_Array

func searchRange(nums []int, target int) []int {
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
