package _03_Binary_Tree_Zigzag_Level_Order_Traversal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_zigzagLevelOrderTest(t *testing.T) {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for {
		fmt.Println(test[len(test)-1])
		test = test[:len(test)-1]
		if len(test) == 0 {
			break
		}
	}
}
func Test_zigzagLevelOrder(t *testing.T) {
	assert.EqualValues(t, [][]int{{1}, {3, 2}, {4, 5}}, zigzagLevelOrder(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}))
}

func Test_zigzagLevelOrder1(t *testing.T) {
	assert.EqualValues(t, [][]int{{1}, {3, 2}, {4, 5}}, zigzagLevelOrder1(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}))
}
