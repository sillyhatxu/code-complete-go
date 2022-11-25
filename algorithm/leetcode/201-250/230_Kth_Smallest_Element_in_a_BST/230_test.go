package _30_Kth_Smallest_Element_in_a_BST

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	input1 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}
	assert.EqualValues(t, 1, kthSmallest(input1, 1))
	assert.EqualValues(t, 2, kthSmallest(input1, 2))
	assert.EqualValues(t, 3, kthSmallest(input1, 3))
	assert.EqualValues(t, 4, kthSmallest(input1, 4))
	assert.EqualValues(t, 5, kthSmallest(input1, 5))
}
