package _46_Permutations

func permute(nums []int) [][]int {
	var result [][]int
	dfs(&result, nums, []int{})
	return result
}

func dfs(result *[][]int, currentNums []int, candidates []int) {
	if len(currentNums) == 0 {
		*result = append(*result, candidates)
		return
	}
	for i := 0; i < len(currentNums); i++ {
		temp := currentNums[i]
		dfs(result, append(append([]int{}, currentNums[0:i]...), currentNums[i+1:]...), append(candidates, temp))
	}
}
