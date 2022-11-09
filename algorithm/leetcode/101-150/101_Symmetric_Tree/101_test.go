package _01_Symmetric_Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	assert.EqualValues(t, true, isSymmetricIteration(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}))
}
