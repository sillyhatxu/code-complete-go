package interview

import (
	"sort"
)

/**
在这个解决方案中，我们首先定义了一个 minDistance 函数，它接受一个表示停车场中每个车位位置的整数数组 park 和一个表示有盖顶的车辆数量的整数 k，并返回最小距离。

在 minDistance 函数中，我们首先对停车场中的车位位置进行排序，然后定义左边界 left 为 1，右边界 right 为最左侧和最右侧车位之间的距离。我们使用二分查找算法，在左边界小于右边界的情况下，不断缩小 left 和 right 的范围，直到找到最小距离。

在每次二分查找中，我们将中间值 mid 定义为 left 和 right 的平均值。然后，我们调用 canPark 函数，判断在距离为 mid 的情况下，是否能够停放 k 辆有盖顶的车辆。如果可以，我们将 right 缩小到 mid，否则将 left 增加到 mid + 1。

在 canPark 函数中，我们首先将停车场中的车位位置进行排序，并初始化计数器 count 为 1，表示第一辆有盖顶的车辆已经停放。然后，我们使用一个循环遍历停车场中的车位位置，如果相邻两个车位之间的距离大于等于 dist，则表示可以在该车位停放一辆有盖顶的车辆，并将计数器 count 加 1。如果 count 大于 k，则表示无法在停车场中停放所有的有盖顶的车辆，函数返回 false。如果循环结束后 count 小于等于 k，则表示可以在停车场中停放所有的有盖顶的车辆，函数返回 true。

通过这个解决方案，我们可以在Golang中找到停车场中最小距离的值
*/
func minDistance(park []int, k int) int {
	sort.Ints(park)
	n := len(park)
	left, right := 1, park[n-1]-park[0]
	for left < right {
		mid := left + (right-left)/2
		if canPark(park, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canPark(park []int, k, dist int) bool {
	n := len(park)
	count := 1
	prev := park[0]
	for i := 1; i < n; i++ {
		if park[i]-prev >= dist {
			count++
			prev = park[i]
			if count > k {
				return false
			}
		}
	}
	return true
}
