package main

import "fmt"

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
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	x := -1
	y := len(nums) - 2
	for x != y {
		if nums[y] >= nums[y+1] {
			y--
			continue
		} else {
			x = y
			y = y + 1
			break
		}
	}
	if y == 0 {
		reverse(&nums, 0, len(nums)-1)
		return
	}
	for i := y + 1; i < len(nums); i++ {
		if nums[x] > nums[i] {
			break
		}
		y = i
	}
	nums[x], nums[y] = nums[y], nums[x]
	reverse(&nums, y, len(nums)-1)
}

func reverse(nums *[]int, start int, end int) {
	for start < end {
		(*nums)[start], (*nums)[end] = (*nums)[end], (*nums)[start]
		start++
		end--
	}
}

func main() {
	var array []int
	array = []int{1, 3, 2}
	nextPermutationOptimize(array)
	fmt.Println(array)

	array = []int{1, 2, 7, 4, 3, 1}
	nextPermutationOptimize(array)
	fmt.Println(array)
	fmt.Println("1,3,1,2,4,7")

	//array = []int{1, 2, 3}
	//nextPermutation(array)
	//fmt.Println(array)
}
