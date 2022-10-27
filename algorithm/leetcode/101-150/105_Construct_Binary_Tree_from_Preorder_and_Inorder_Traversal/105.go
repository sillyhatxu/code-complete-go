package _05_Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//TODO
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	result := &TreeNode{
		Val: preorder[0],
	}
	if len(inorder) == 1 {
		return result
	}
	index := getIndexValue(result.Val, inorder) + 1
	result.Left = buildTree(preorder[1:index], inorder[:index])
	result.Right = buildTree(preorder[index:], inorder[index:])
	return result
}

/**
   1
  2  3
   4   5
*/
func getIndexValue(val int, array []int) int {
	for i, v := range array {
		if v == val {
			return i
		}
	}
	return -1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
