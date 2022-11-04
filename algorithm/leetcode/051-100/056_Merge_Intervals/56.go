package _56_Merge_Intervals

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	// Time: O(nlogn)
	// Space: O(1)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// Time: O(n)
	// Space: O(1)
	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] >= intervals[i][1] {
			intervals = append(intervals[:i], intervals[i+1:]...)
			i--
			continue
		} else if intervals[i][0] <= intervals[i-1][1] && intervals[i-1][1] <= intervals[i][1] {
			intervals[i] = []int{intervals[i-1][0], intervals[i][1]}
			intervals = append(intervals[:i-1], intervals[i:]...)
			i--
			continue
		}
	}
	return intervals
}

func merge1(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	// Time: O(nlogn)
	// Space: O(1)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// Time: O(n)
	// Space: O(n)
	var result [][]int
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lenMergeIntervals := len(result)
		if intervals[i][0] <= result[lenMergeIntervals-1][1] {
			if intervals[i][1] > result[lenMergeIntervals-1][1] {
				result[lenMergeIntervals-1][1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}
