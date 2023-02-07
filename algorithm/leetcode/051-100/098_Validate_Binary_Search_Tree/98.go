package _98_Validate_Binary_Search_Tree

import (
	"math"
)

func isValidBST(root *TreeNode) bool {
	var helper func(node *TreeNode, min, max int) bool
	helper = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return helper(node.Left, min, node.Val) && helper(node.Right, node.Val, max)
	}
	return helper(root.Left, math.MinInt64, root.Val) && helper(root.Right, root.Val, math.MaxInt64)
}

func isValidBST2(root *TreeNode) bool {
	previous := math.MinInt64
	_, result := inorderAndCompare(root, &previous)
	return result
}

func inorderAndCompare(node *TreeNode, previous *int) (*int, bool) {
	if node == nil {
		return previous, true
	}
	previous, leftResult := inorderAndCompare(node.Left, previous)
	if !leftResult {
		return previous, false
	}
	if *previous >= node.Val {
		return previous, false
	}
	previous = &node.Val
	previous, rightResult := inorderAndCompare(node.Right, previous)
	if !rightResult {
		return previous, false
	}
	return previous, true
}

func isValidBST1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var trees []int
	dfs(root, &trees)
	for i := 1; i < len(trees); i++ {
		if trees[i-1] >= trees[i] {
			return false
		}
	}
	return true
}

func dfs(node *TreeNode, trees *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, trees)
	*trees = append(*trees, node.Val)
	dfs(node.Right, trees)
}

func isValidBSTIteration(root *TreeNode) bool {
	var previous *int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append([]*TreeNode{root}, stack...)
			root = root.Left
			continue
		}
		root = stack[0]
		if previous != nil && root.Val <= *previous {
			return false
		}
		previous = &root.Val
		root = root.Right
		stack = stack[1:]
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
