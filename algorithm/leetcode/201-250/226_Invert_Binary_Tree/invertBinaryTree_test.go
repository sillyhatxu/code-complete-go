package invertBinaryTree

import (
	"fmt"
	"testing"
)

func Test_invertTree(t *testing.T) {
	test1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	result := invertTree(test1)
	preOrderTraversal(result)
	fmt.Println()
}

func preOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Val, ", ")
	preOrderTraversal(node.Left)
	preOrderTraversal(node.Right)
}
