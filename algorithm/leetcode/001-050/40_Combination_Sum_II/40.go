package _0_Combination_Sum_II

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var dfs func(result *[][]int, candidates []int, target int, selection *[]int, index int)
	dfs = func(result *[][]int, candidates []int, target int, selection *[]int, index int) {
		if target < 0 {
			return
		} else if target == 0 {
			temp := make([]int, len(*selection), len(*selection))
			copy(temp, *selection)
			*result = append(*result, temp)
			return
		}
		for i := index; i < len(candidates); i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			tempTarget := target - candidates[i]
			*selection = append(*selection, candidates[i])
			dfs(result, candidates, tempTarget, selection, i+1)
			*selection = (*selection)[:len(*selection)-1]
		}
	}
	var result [][]int
	sort.Ints(candidates)
	dfs(&result, candidates, target, &[]int{}, 0)
	return result
}
