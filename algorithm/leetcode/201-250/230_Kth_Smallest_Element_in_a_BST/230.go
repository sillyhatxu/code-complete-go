package _30_Kth_Smallest_Element_in_a_BST

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var arr []*TreeNode
	for {
		for root != nil {
			arr = append(arr, root)
			root = root.Left
		}
		root = arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
