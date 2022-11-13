package _28_Longest_Consecutive_Sequence

func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool, len(nums))
	max := 0 //max consecutive number
	for _, num := range nums {
		numMap[num] = true
	}
	for num := range numMap {
		if numMap[num-1] {
			continue //not start element
		}
		current := num
		for numMap[current+1] {
			current++
		}
		curCon := current - num + 1 //current consecutive number
		//get max
		if curCon > max {
			max = curCon
		}
	}
	return max
}
