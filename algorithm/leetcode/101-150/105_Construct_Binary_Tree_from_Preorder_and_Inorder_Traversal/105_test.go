package _05_Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildTree1(t *testing.T) {
	result := buildTree([]int{1, 2, 3, 4, 5, 6, 7}, []int{7, 6, 5, 4, 3, 2, 1})
	fmt.Println(result)
}
func Test_buildTree(t *testing.T) {
	result := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	assert.EqualValues(t, 3, result.Val)
	assert.EqualValues(t, 9, result.Left.Val)
	assert.Nil(t, result.Left.Left)
	assert.Nil(t, result.Left.Right)
	assert.EqualValues(t, 20, result.Right.Val)
	assert.EqualValues(t, 15, result.Right.Left.Val)
	assert.EqualValues(t, 7, result.Right.Right.Val)
}
