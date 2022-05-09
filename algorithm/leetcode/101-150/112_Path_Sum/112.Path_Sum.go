package main

import "fmt"

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

func main() {
	mockData := []*int{
		toIntPoint(5),
		toIntPoint(4),
		toIntPoint(8),
		toIntPoint(11),
		nil,
		toIntPoint(13),
		toIntPoint(4),
		toIntPoint(7),
		toIntPoint(2),
		nil, nil, nil,
		toIntPoint(1),
	}
	binaryTree := buildBinaryTree(mockData)
	fmt.Println(hasPathSum(binaryTree, 22))
}

func buildBinaryTree(input []*int) *TreeNode {
	if input[0] == nil {
		return nil
	}
	treeNode := &TreeNode{}
	return createTreeNode(input, treeNode, 0)
}

func createTreeNode(input []*int, treeNode *TreeNode, index int) *TreeNode {
	if index >= len(input) {
		return nil
	}
	treeNode.Val = *input[index]
	treeNode.Left = createTreeNode(input, treeNode.Left, 2*index+1)
	treeNode.Right = createTreeNode(input, treeNode.Left, 2*index+2)
	return treeNode
}

func toIntPoint(input int) *int {
	return &input
}
