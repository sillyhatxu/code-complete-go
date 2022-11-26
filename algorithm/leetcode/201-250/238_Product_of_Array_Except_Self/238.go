package _38_Product_of_Array_Except_Self

func productExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length, length)
	res[0] = 1
	current := nums[0]
	for i := 1; i < length; i++ {
		res[i] = current
		current *= nums[i]
	}
	current = nums[length-1]
	for i := length - 2; i >= 0; i-- {
		res[i] *= current
		current *= nums[i]
	}
	return res
}
