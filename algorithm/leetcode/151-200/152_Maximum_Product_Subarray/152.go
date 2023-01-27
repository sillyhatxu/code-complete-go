package _52_Maximum_Product_Subarray

func maxProduct(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	res, maxValue, minValue := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tempMax := maxValue
		maxValue = max(maxValue*nums[i], max(minValue*nums[i], nums[i]))
		minValue = min(tempMax*nums[i], min(minValue*nums[i], nums[i]))
		res = max(res, maxValue)
	}
	return res
}

func maxProduct1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	max := nums[0]
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		tempMax := max
		max = getMax(max*nums[i], getMax(min*nums[i], nums[i]))
		min = getMin(tempMax*nums[i], getMin(min*nums[i], nums[i]))
		res = getMax(res, max)
	}
	return res
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func getMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
