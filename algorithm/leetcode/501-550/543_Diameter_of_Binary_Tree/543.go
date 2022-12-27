package _43_Diameter_of_Binary_Tree

func diameterOfBinaryTree(root *TreeNode) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
	}
	var diameter func(root *TreeNode) int
	diameter = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return max(maxDepth(root.Left)+maxDepth(root.Right), max(diameter(root.Left), diameter(root.Right)))
	}
	return diameter(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
