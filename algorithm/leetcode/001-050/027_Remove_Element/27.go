package main

func removeElement1(nums []int, val int) int {
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
