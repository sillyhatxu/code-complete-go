package _60_Subarray_Sum_Equals_K

func subarraySum(nums []int, k int) int {
	res, sum := 0, 0
	validate := make(map[int]int)
	validate[0] = 1
	for _, num := range nums {
		sum += num
		key := sum - k
		if v, ok := validate[key]; ok {
			res += v
		}
		if v, ok := validate[sum]; ok {
			validate[sum] = v + 1
		} else {
			validate[sum] = 1
		}
	}
	return res
}
