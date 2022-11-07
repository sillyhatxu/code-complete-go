package _88_Merge_Sorted_Array

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	test1 := func() []int {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		nums2 := []int{2, 5, 6}
		merge(nums1, 3, nums2, 3)
		return nums1
	}
	fmt.Println(test1())

	test2 := func() []int {
		nums1 := []int{0}
		nums2 := []int{1}
		merge(nums1, 0, nums2, 1)
		return nums1
	}
	fmt.Println(test2())

	test3 := func() []int {
		nums1 := []int{1, 0}
		nums2 := []int{2}
		merge(nums1, 1, nums2, 1)
		return nums1
	}
	fmt.Println(test3())
}
