package _37_Path_Sum_III

func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	var helper func(node *TreeNode, target int)
	helper = func(node *TreeNode, target int) {
		if node == nil {
			return
		}
		if node.Val == target {
			res++
		}
		nextTarget := target - node.Val
		helper(node.Left, nextTarget)
		helper(node.Right, nextTarget)
	}
	var sum func(root *TreeNode, target int)
	sum = func(node *TreeNode, target int) {
		if node == nil {
			return
		}
		helper(node, target)
		sum(node.Left, target)
		sum(node.Right, target)
	}
	sum(root, targetSum)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
