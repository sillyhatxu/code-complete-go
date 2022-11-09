package _98_Validate_Binary_Search_Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	assert.EqualValues(t, false, isValidBST(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
	}))
	assert.EqualValues(t, true, isValidBST(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
	assert.EqualValues(t, false, isValidBST(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}))
	assert.EqualValues(t, false, isValidBST(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}))
	assert.EqualValues(t, true, isValidBST(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}))
}
