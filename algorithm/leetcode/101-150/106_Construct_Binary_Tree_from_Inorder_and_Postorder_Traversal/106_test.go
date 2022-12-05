package _06_Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildTree1(t *testing.T) {
	result := buildTree([]int{3, 2, 1}, []int{3, 2, 1})
	assert.EqualValues(t, 1, result.Val)
	assert.EqualValues(t, 2, result.Left.Val)
	assert.EqualValues(t, 3, result.Left.Left.Val)
}
func Test_buildTree(t *testing.T) {
	result := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	assert.EqualValues(t, 3, result.Val)
	assert.EqualValues(t, 9, result.Left.Val)
	assert.Nil(t, result.Left.Left)
	assert.Nil(t, result.Left.Right)
	assert.EqualValues(t, 20, result.Right.Val)
	assert.EqualValues(t, 15, result.Right.Left.Val)
	assert.EqualValues(t, 7, result.Right.Right.Val)
}
