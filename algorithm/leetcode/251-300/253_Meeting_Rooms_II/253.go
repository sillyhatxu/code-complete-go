package _53_Meeting_Rooms_II

import "sort"

func minMeetingRooms(intervals [][]int) int {
	nums := make([]int, len(intervals)*2, len(intervals)*2)
	for i, interval := range intervals {
		nums[i*2] = interval[0]*10 + 2
		nums[i*2+1] = interval[1]*10 + 1
	}
	sort.Ints(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxRoom, currentRoom := 0, 0
	for _, num := range nums {
		if num%10 == 2 {
			currentRoom++
		} else {
			//intervals[i][0] %10 == 2
			currentRoom--
		}
		maxRoom = max(currentRoom, maxRoom)
	}
	return maxRoom
}
