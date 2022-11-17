package _52_Maximum_Product_Subarray

func maxProduct(nums []int) int {
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
