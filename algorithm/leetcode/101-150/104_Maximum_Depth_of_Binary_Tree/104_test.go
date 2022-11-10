package _04_Maximum_Depth_of_Binary_Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	assert.EqualValues(t, 3, maxDepth(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
	assert.EqualValues(t, 3, maxDepth(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}))
}
