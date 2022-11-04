# 56. Merge Intervals

Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

### Example 1:

```
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
```

### Example 2:

```
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping
```

### Constraints:

* 1 <= intervals.length <= 104
* intervals[i].length == 2
* 0 <= starti <= endi <= 104

### Translate:

> 56. 合并区间
> 
> 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
>
> 示例 1：
> 
> 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
> 
> 输出：[[1,6],[8,10],[15,18]]
> 
> 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
> 
> 示例2：
> 
> 输入：intervals = [[1,4],[4,5]]
> 
> 输出：[[1,5]]
> 
> 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

### 解题思路
起始位置排序后，只有两种情况需要合并
1. [1,4],[2,3] 和 [1,4],[2,4]
2. [1,4],[2,8]

因此直接遍历即可。

但是操作数组时，如果使用
空间复杂度O(1)的方式，需要对原数组做移除元素操作，所以运行效率会变慢
如果新增一个数组，空间复杂度变为O(n)
### 代码

#### Space: O(1)

```golang
ffunc merge(intervals [][]int) [][]int {
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
    for i:=1; i < len(intervals);i++{
        if intervals[i-1][1] >= intervals[i][1]{
            intervals = append(intervals[:i],intervals[i+1:]...)
            i--
            continue
        }else if  intervals[i][0] <= intervals[i-1][1] && intervals[i-1][1] <= intervals[i][1] {
            intervals[i] = []int{intervals[i-1][0],intervals[i][1]}
            intervals = append(intervals[:i-1],intervals[i:]...)
            i--
            continue
        }
    }
    return intervals
}
```

#### Space: O(n)

```golang
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
	// Space: O(n)
	var result [][]int
	result = append(result,intervals[0])
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
```