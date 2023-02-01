package _02_Binary_Tree_Level_Order_Traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	assert.EqualValues(t, [][]int{{1}, {2}, {3}, {4}, {5}}, levelOrder(input1))
	assert.EqualValues(t, [][]int{{3}, {9, 20}, {15, 17}}, levelOrder(input2))
}

func Test_levelOrder1(t *testing.T) {
	assert.EqualValues(t, [][]int{{1}, {2}, {3}, {4}, {5}}, levelOrder1(input1))
	assert.EqualValues(t, [][]int{{3}, {9, 20}, {15, 17}}, levelOrder1(input2))
}

var input1 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
			},
		},
	},
}

var input2 = &TreeNode{
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
			Val: 17,
		},
	},
}
