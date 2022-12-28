package _14_Flatten_Binary_Tree_to_Linked_List

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var stack []*TreeNode
	stack = append([]*TreeNode{root}, stack...)
	for len(stack) != 0 {
		curr := stack[0]
		stack = stack[1:]
		if curr.Right != nil {
			stack = append([]*TreeNode{curr.Right}, stack...)
		}
		if curr.Left != nil {
			stack = append([]*TreeNode{curr.Left}, stack...)
		}
		if len(stack) != 0 {
			curr.Right = stack[0]
		}
		curr.Left = nil
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
