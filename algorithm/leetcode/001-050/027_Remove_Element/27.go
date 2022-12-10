package main

func removeElement(nums []int, val int) int {
	cur, res := 0, len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			res--
			continue
		}
		nums[cur] = nums[i]
		cur++
	}
	return res
}
