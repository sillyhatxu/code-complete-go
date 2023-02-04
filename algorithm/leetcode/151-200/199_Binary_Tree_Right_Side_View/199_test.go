package _99_Binary_Tree_Right_Side_View

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	assert.EqualValues(t, []int{1, 3, 4}, rightSideView(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
	}))
	assert.EqualValues(t, []int{1, 3, 6}, rightSideView(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}},
		Right: &TreeNode{Val: 3},
	}))
}

func Test_rightSideView1(t *testing.T) {
	assert.EqualValues(t, []int{1, 3, 4}, rightSideView1(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
	}))
	assert.EqualValues(t, []int{1, 3, 6}, rightSideView1(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}},
		Right: &TreeNode{Val: 3},
	}))
}
