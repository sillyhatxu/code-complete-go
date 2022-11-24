package _17_Contains_Duplicate

func containsDuplicate(nums []int) bool {
	check := make(map[int]struct{})
	for i := range nums {
		if _, ok := check[nums[i]]; ok {
			return true
		}
		check[nums[i]] = struct{}{}
	}
	return false
}
