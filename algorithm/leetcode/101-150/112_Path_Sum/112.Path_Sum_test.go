package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	assert.EqualValues(t, true, hasPathSum(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}, 22))
	assert.EqualValues(t, false, hasPathSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}, 5))
	assert.EqualValues(t, true, hasPathSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}, 4))
	assert.EqualValues(t, false, hasPathSum(nil, 4))
}
