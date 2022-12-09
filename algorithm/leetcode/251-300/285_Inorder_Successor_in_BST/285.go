package _85_Inorder_Successor_in_BST

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return nil
	}
	var res *TreeNode
	var pre *TreeNode
	var dfs func(currentNode *TreeNode, target *TreeNode)
	dfs = func(currentNode *TreeNode, target *TreeNode) {
		if currentNode == nil {
			return
		}
		dfs(currentNode.Left, target)
		if pre == target {
			res, pre = currentNode, currentNode
			return
		}
		pre = currentNode
		dfs(currentNode.Right, target)
	}
	dfs(root, p)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
