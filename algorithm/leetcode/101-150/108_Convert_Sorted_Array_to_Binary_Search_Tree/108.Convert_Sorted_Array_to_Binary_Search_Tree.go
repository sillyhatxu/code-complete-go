package main

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(&nums, 0, len(nums)-1)
}

func helper(nums *[]int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	return &TreeNode{
		Val:   (*nums)[mid],
		Left:  helper(nums, left, mid-1),
		Right: helper(nums, mid+1, right),
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
