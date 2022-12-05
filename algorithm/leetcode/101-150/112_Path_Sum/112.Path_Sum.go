package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	balance := targetSum - root.Val
	if root.Left == nil && root.Right == nil {
		return balance == 0
	}
	return hasPathSum(root.Left, balance) || hasPathSum(root.Right, balance)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
