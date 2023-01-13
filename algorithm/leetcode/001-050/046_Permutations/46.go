package _46_Permutations

func permute(nums []int) [][]int {
	var dfs func(res *[][]int, currents, candidates []int)
	dfs = func(res *[][]int, currents, candidates []int) {
		if len(currents) == 0 {
			*res = append(*res, candidates)
			return
		}
		for i := 0; i < len(currents); i++ {
			tempCurrents := append(append([]int{}, currents[0:i]...), currents[i+1:]...)
			tempCandidates := append(candidates, currents[i])
			dfs(res, tempCurrents, tempCandidates)
		}
	}
	var res [][]int
	dfs(&res, nums, []int{})
	return res
}

func permute1(nums []int) [][]int {
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
