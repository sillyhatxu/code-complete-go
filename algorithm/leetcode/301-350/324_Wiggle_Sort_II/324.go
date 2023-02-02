package _24_Wiggle_Sort_II

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) {
	sort.Ints(nums)
	temp := make([]int, len(nums), len(nums))
	copy(temp, nums)
	mid := (len(nums)+1)/2 - 1
	last := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			nums[i] = temp[mid]
			mid--
		} else {
			nums[i] = temp[last]
			last--
		}
	}
}

//Time: O(n); Space: O(n)
func wiggleSort1(nums []int) {
	sort.Ints(nums)
	temp := make([]int, len(nums), len(nums))
	for i := range nums {
		temp[i] = nums[i]
	}
	midIndex := (len(nums)+1)/2 - 1
	endIndex := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			nums[i] = temp[midIndex]
			midIndex--
		} else {
			nums[i] = temp[endIndex]
			endIndex--
		}
	}
}

//Time: O(n); Space: O(1)
func wiggleSort2(nums []int) {
	length := len(nums)
	mid := (length+1)/2 - 1
	largePos := 1 //Odd position from start for larger than median numbers
	smallPos := 0 //Even position from end for smaller than median numbers
	if length%2 == 0 {
		smallPos = length - 2
	} else {
		smallPos = length - 1
	}
	cur := 0
	for cur < length {
		if nums[cur] < mid && (cur < smallPos || (cur >= smallPos && cur%2 != 0)) {
			//Avoiding already checked even positions from the end
			nums[cur], nums[smallPos] = nums[smallPos], nums[cur]
			smallPos -= 2
		} else if nums[cur] > mid && (largePos < cur || (largePos >= cur && cur%2 == 0)) {
			//Avoiding already checked odd positions from the start
			nums[cur], nums[largePos] = nums[largePos], nums[cur]
			largePos += 2
		} else {
			cur++
		}
	}
	fmt.Println(nums)
}
