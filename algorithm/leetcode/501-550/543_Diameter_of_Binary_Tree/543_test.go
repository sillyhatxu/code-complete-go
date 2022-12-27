package _43_Diameter_of_Binary_Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	assert.EqualValues(t, 3, diameterOfBinaryTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
}
