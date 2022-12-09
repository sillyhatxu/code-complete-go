package _85_Inorder_Successor_in_BST

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_inorderSuccessor(t *testing.T) {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node5.Left = node3
	node5.Right = node6
	node3.Left = node2
	node3.Right = node4
	node2.Left = node1
	assert.Nil(t, inorderSuccessor(node5, node6))
	assert.EqualValues(t, node3, inorderSuccessor(node5, node2))
	assert.EqualValues(t, node4, inorderSuccessor(node5, node3))
	assert.EqualValues(t, node5, inorderSuccessor(node5, node4))
}
