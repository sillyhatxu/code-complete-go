package _06_Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildTree(t *testing.T) {
	result1 := buildTree([]int{3, 2, 1}, []int{3, 2, 1})
	assert.EqualValues(t, 1, result1.Val)
	assert.EqualValues(t, 2, result1.Left.Val)
	assert.EqualValues(t, 3, result1.Left.Left.Val)

	result2 := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	assert.EqualValues(t, 3, result2.Val)
	assert.EqualValues(t, 9, result2.Left.Val)
	assert.Nil(t, result2.Left.Left)
	assert.Nil(t, result2.Left.Right)
	assert.EqualValues(t, 20, result2.Right.Val)
	assert.EqualValues(t, 15, result2.Right.Left.Val)
	assert.EqualValues(t, 7, result2.Right.Right.Val)
}

func Test_buildTree1(t *testing.T) {
	result1 := buildTree1([]int{3, 2, 1}, []int{3, 2, 1})
	assert.EqualValues(t, 1, result1.Val)
	assert.EqualValues(t, 2, result1.Left.Val)
	assert.EqualValues(t, 3, result1.Left.Left.Val)

	result2 := buildTree1([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	assert.EqualValues(t, 3, result2.Val)
	assert.EqualValues(t, 9, result2.Left.Val)
	assert.Nil(t, result2.Left.Left)
	assert.Nil(t, result2.Left.Right)
	assert.EqualValues(t, 20, result2.Right.Val)
	assert.EqualValues(t, 15, result2.Right.Left.Val)
	assert.EqualValues(t, 7, result2.Right.Right.Val)
}
