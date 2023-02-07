package _98_Validate_Binary_Search_Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	assert.EqualValues(t, true, isValidBST(input6))
	assert.EqualValues(t, false, isValidBST(input1))
	assert.EqualValues(t, true, isValidBST(input2))
	assert.EqualValues(t, false, isValidBST(input3))
	assert.EqualValues(t, false, isValidBST(input4))
	assert.EqualValues(t, true, isValidBST(input5))
}

func Test_isValidBST1(t *testing.T) {
	assert.EqualValues(t, false, isValidBST1(input1))
	assert.EqualValues(t, true, isValidBST1(input2))
	assert.EqualValues(t, false, isValidBST1(input3))
	assert.EqualValues(t, false, isValidBST1(input4))
	assert.EqualValues(t, true, isValidBST1(input5))
}

func Test_isValidBST2(t *testing.T) {
	assert.EqualValues(t, false, isValidBST2(input1))
	assert.EqualValues(t, true, isValidBST2(input2))
	assert.EqualValues(t, false, isValidBST2(input3))
	assert.EqualValues(t, false, isValidBST2(input4))
	assert.EqualValues(t, true, isValidBST2(input5))
}

var input1 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 1,
	},
}
var input2 = &TreeNode{
	Val: 2,
	Left: &TreeNode{
		Val: 1,
	},
	Right: &TreeNode{
		Val: 3,
	},
}

var input3 = &TreeNode{
	Val: 2,
	Left: &TreeNode{
		Val: 2,
	},
	Right: &TreeNode{
		Val: 2,
	},
}

var input4 = &TreeNode{
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
}

var input5 = &TreeNode{
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
}

var input6 = &TreeNode{
	Val: -2147483648,
	Right: &TreeNode{
		Val: 2147483647,
	},
}
