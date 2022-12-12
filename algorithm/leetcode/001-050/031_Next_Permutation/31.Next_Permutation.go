package main

func nextPermutation(nums []int) {
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := len(nums) - 1; j > i; j-- {
				if nums[j] > nums[i] {
					nums[i], nums[j] = nums[j], nums[i]
					return
				}
			}
		}
	}
}

func nextPermutationOptimize(nums []int) {
	reverse := func(nums *[]int, start int, end int) {
		for start < end {
			(*nums)[start], (*nums)[end] = (*nums)[end], (*nums)[start]
			start++
			end--
		}
	}
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	x := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			x = i
			break
		}
	}
	if x == -1 {
		reverse(&nums, 0, len(nums)-1)
		return
	}
	y := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[x] {
			y = i
			break
		}
	}
	nums[x], nums[y] = nums[y], nums[x]
	reverse(&nums, x+1, len(nums)-1)
}
