package _44_Binary_Tree_Preorder_Traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_preOrderTraversal(t *testing.T) {
	result := preorderTraversal(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	})
	assert.EqualValues(t, 1, result[0])
	assert.EqualValues(t, 2, result[1])
	assert.EqualValues(t, 3, result[2])
}

func Test_preOrderTraversal1(t *testing.T) {
	result := preorderTraversalOptimal(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	})
	assert.EqualValues(t, 1, result[0])
	assert.EqualValues(t, 2, result[1])
	assert.EqualValues(t, 3, result[2])
}
