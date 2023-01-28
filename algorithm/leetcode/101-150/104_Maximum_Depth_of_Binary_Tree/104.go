package _04_Maximum_Depth_of_Binary_Tree

func maxDepth(root *TreeNode) int {
	var recursion func(node *TreeNode, currentDepth int) int
	recursion = func(node *TreeNode, currentDepth int) int {
		if node == nil {
			return currentDepth
		}
		currentDepth++
		leftDepth := recursion(node.Left, currentDepth)
		rightDepth := recursion(node.Right, currentDepth)
		if leftDepth > rightDepth {
			return leftDepth
		} else {
			return rightDepth
		}
	}
	return recursion(root, 0)
}

func maxDepth2(root *TreeNode) int {
	depth := 0
	var dfs func(node *TreeNode, currentDepth int)
	dfs = func(node *TreeNode, currentDepth int) {
		if node == nil {
			return
		}
		currentDepth++
		if currentDepth > depth {
			depth = currentDepth
		}
		dfs(node.Left, currentDepth)
		dfs(node.Right, currentDepth)
	}
	dfs(root, 0)
	return depth
}

func maxDepth1(root *TreeNode) int {
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
