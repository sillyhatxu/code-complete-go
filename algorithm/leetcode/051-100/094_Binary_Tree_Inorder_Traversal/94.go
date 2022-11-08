package _94_Binary_Tree_Inorder_Traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	var stack []*TreeNode
	//currentNote := root
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append([]*TreeNode{root}, stack...)
			root = root.Left
			continue
		}
		result = append(result, stack[0].Val)
		root = stack[0].Right
		stack = stack[1:]

	}
	return result
}

func inorderTraversalRecursion(root *TreeNode) []int {
	var result []int
	dfs(&result, root)
	return result
}

func dfs(result *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	dfs(result, node.Left)
	*result = append(*result, node.Val)
	dfs(result, node.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
