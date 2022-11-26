package _36_Lowest_Common_Ancestor_of_a_Binary_Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	treeNode0 := &TreeNode{Val: 0}
	treeNode1 := &TreeNode{Val: 1}
	treeNode2 := &TreeNode{Val: 2}
	treeNode3 := &TreeNode{Val: 3}
	treeNode4 := &TreeNode{Val: 4}
	treeNode5 := &TreeNode{Val: 5}
	treeNode6 := &TreeNode{Val: 6}
	treeNode7 := &TreeNode{Val: 7}
	treeNode8 := &TreeNode{Val: 8}
	treeNode3.Left, treeNode3.Right = treeNode5, treeNode1
	treeNode5.Left, treeNode5.Right = treeNode6, treeNode2
	treeNode2.Left, treeNode2.Right = treeNode7, treeNode4
	treeNode1.Left, treeNode1.Right = treeNode0, treeNode8
	//assert.EqualValues(t, treeNode3, lowestCommonAncestor(treeNode3, treeNode5, treeNode1))
	//assert.EqualValues(t, treeNode1, lowestCommonAncestor(treeNode3, treeNode0, treeNode8))
	//assert.EqualValues(t, treeNode5, lowestCommonAncestor(treeNode3, treeNode5, treeNode4))
	//assert.EqualValues(t, treeNode2, lowestCommonAncestor(treeNode3, treeNode2, treeNode7))
	assert.EqualValues(t, treeNode3, lowestCommonAncestor(treeNode3, treeNode6, treeNode8))
}
