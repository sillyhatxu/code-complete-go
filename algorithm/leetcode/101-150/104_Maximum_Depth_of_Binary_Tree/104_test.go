package _04_Maximum_Depth_of_Binary_Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	assert.EqualValues(t, 3, maxDepth(test1))
	assert.EqualValues(t, 3, maxDepth(test2))
}

func Test_maxDepth1(t *testing.T) {
	assert.EqualValues(t, 3, maxDepth1(test1))
	assert.EqualValues(t, 3, maxDepth1(test2))
}

func Test_maxDepth2(t *testing.T) {
	assert.EqualValues(t, 3, maxDepth2(test1))
	assert.EqualValues(t, 3, maxDepth2(test2))
}

var test1 = &TreeNode{
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
}

var test2 = &TreeNode{
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
}
