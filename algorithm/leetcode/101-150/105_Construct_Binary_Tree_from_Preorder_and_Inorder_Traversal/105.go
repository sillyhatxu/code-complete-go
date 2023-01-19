package _05_Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderIndexMap := make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		inorderIndexMap[inorder[i]] = i
	}
	var helper func(preorder []int, inorder []int, preIndex, inStart, inEnd int, inorderIndexMap map[int]int) *TreeNode
	helper = func(preorder []int, inorder []int, preIndex, inStart, inEnd int, inorderIndexMap map[int]int) *TreeNode {
		if inStart > inEnd {
			return nil
		}
		rootIndex := inorderIndexMap[preorder[preIndex]]
		return &TreeNode{
			Val:   preorder[preIndex],
			Left:  helper(preorder, inorder, preIndex+1, inStart, rootIndex-1, inorderIndexMap),
			Right: helper(preorder, inorder, preIndex+(rootIndex-inStart+1), rootIndex+1, inEnd, inorderIndexMap),
		}
	}
	return helper(preorder, inorder, 0, 0, len(inorder)-1, inorderIndexMap)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree2(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		inorderMap[inorder[i]] = i
	}
	return helper(preorder, inorder, 0, 0, len(inorder)-1, inorderMap)
}

func helper(preorder []int, inorder []int, preIndex int, inStart, inEnd int, inorderMap map[int]int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	root := &TreeNode{Val: preorder[preIndex]}
	rootIndex := inorderMap[preorder[preIndex]]
	root.Left = helper(preorder, inorder, preIndex+1, inStart, rootIndex-1, inorderMap)
	rightPreIndex := preIndex + (rootIndex - inStart + 1) //(rootIndex - inStart + 1): the total number of left node
	root.Right = helper(preorder, inorder, rightPreIndex, rootIndex+1, inEnd, inorderMap)
	return root
}

func buildTree1(preorder []int, inorder []int) *TreeNode {
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
	result.Left = buildTree1(preorder[1:index], inorder[:index])
	result.Right = buildTree1(preorder[index:], inorder[index:])
	return result
}
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
