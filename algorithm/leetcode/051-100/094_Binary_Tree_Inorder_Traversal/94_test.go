package _94_Binary_Tree_Inorder_Traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	assert.EqualValues(t, []int{1, 3, 2}, inorderTraversal(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}))
}

func Test_inorderTraversalRecursion(t *testing.T) {
	assert.EqualValues(t, []int{1, 3, 2}, inorderTraversalRecursion(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}))
}
