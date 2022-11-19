package _69_Majority_Element

func majorityElement(nums []int) int {
	count := 1
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			count++
			result = nums[i]
			continue
		}
		if nums[i] == result {
			count++
		} else {
			count--
		}
	}
	return result
}

func majorityElement1(nums []int) int {
	count := 1
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			result = nums[i]
		}
		if nums[i] == result {
			count++
		} else {
			count--
		}
	}
	return result
}
