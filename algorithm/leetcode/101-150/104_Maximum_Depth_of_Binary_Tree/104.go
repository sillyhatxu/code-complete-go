package _04_Maximum_Depth_of_Binary_Tree

func maxDepth(root *TreeNode) int {
	depth := 0
	dfs(root, &depth, 0)
	return depth
}

func dfs(node *TreeNode, depth *int, currentDepth int) {
	if node == nil {
		return
	}
	currentDepth++
	if *depth < currentDepth {
		*depth = currentDepth
	}
	dfs(node.Left, depth, currentDepth)
	dfs(node.Right, depth, currentDepth)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
