package _14_Flatten_Binary_Tree_to_Linked_List

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_flatten(t *testing.T) {
	root := &TreeNode{
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
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	flatten(root)
	assert.EqualValues(t, 1, root.Val)
	assert.EqualValues(t, 2, root.Right.Val)
	assert.EqualValues(t, 3, root.Right.Right.Val)
	assert.EqualValues(t, 4, root.Right.Right.Right.Val)
	assert.EqualValues(t, 5, root.Right.Right.Right.Right.Val)
	assert.EqualValues(t, 6, root.Right.Right.Right.Right.Right.Val)

}
