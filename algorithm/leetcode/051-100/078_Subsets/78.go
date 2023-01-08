package _78_Subsets

func subsets(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	var helper func(res *[][]int, nums []int, curr []int, index int)
	helper = func(res *[][]int, nums []int, curr []int, index int) {
		cp := make([]int, len(curr), len(curr))
		copy(cp, curr)
		*res = append(*res, cp)
		for i := index; i < len(nums); i++ {
			curr = append(curr, nums[i])
			helper(res, nums, curr, i+1)
			curr = curr[:len(curr)-1]
		}
	}
	helper(&res, nums, []int{}, 0)
	return res
}

func subsets1(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		return result
	}
	helper(&result, []int{}, nums, 0)
	return result
}

func helper(result *[][]int, temp []int, nums []int, index int) {
	copyTemp := make([]int, len(temp))
	copy(copyTemp, temp)
	*result = append(*result, copyTemp)
	for i := index; i < len(nums); i++ {
		temp = append(temp, nums[i])
		helper(result, temp, nums, i+1)
		temp = temp[:len(temp)-1]
	}
}
