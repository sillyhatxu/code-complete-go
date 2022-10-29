package _200_Minimum_Absolute_Difference

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	result := [][]int{{arr[0], arr[1]}}
	minimum := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		tempMin := arr[i+1] - arr[i]
		if tempMin < minimum {
			result = nil
			minimum = tempMin
		}
		if minimum == tempMin {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}
	return result
}
