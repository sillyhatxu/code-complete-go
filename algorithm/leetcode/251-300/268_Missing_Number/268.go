package _68_Missing_Number

func missingNumber(nums []int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res ^= nums[i] ^ i
	}
	return res
}
