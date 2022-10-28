package _44_Binary_Tree_Preorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int

	result = append(result, root.Val)

	leftResult := preorderTraversal(root.Left)
	if leftResult != nil {
		result = append(result, leftResult...)
	}
	rightResult := preorderTraversal(root.Right)
	if rightResult != nil {
		result = append(result, rightResult...)
	}
	return result
}

func preorderTraversalOptimal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]
		result = append(result, current.Val)
		if current.Right != nil {
			stack = append([]*TreeNode{current.Right}, stack...)
		}
		if current.Left != nil {
			stack = append([]*TreeNode{current.Left}, stack...)
		}
	}
	return result
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]
		result = append([]int{current.Val}, result...)
		if current.Left != nil {
			stack = append([]*TreeNode{current.Left}, stack...)
		}
		if current.Right != nil {
			stack = append([]*TreeNode{current.Right}, stack...)
		}
	}
	return result
}
