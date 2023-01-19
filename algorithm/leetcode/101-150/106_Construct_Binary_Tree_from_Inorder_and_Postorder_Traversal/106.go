package _06_Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal

func buildTree(inorder []int, postorder []int) *TreeNode {
	inordeIndexrMap := make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		inordeIndexrMap[inorder[i]] = i
	}
	var helper func(inorder []int, postorder []int, postIndex, inStart, inEnd int, inordeIndexrMap map[int]int) *TreeNode
	helper = func(inorder []int, postorder []int, postIndex, inStart, inEnd int, inordeIndexrMap map[int]int) *TreeNode {
		if inStart > inEnd {
			return nil
		}
		rootIndex := inordeIndexrMap[postorder[postIndex]]
		return &TreeNode{
			Val:   postorder[postIndex],
			Left:  helper(inorder, postorder, postIndex-(inEnd-rootIndex)-1, inStart, rootIndex-1, inordeIndexrMap),
			Right: helper(inorder, postorder, postIndex-1, rootIndex+1, inEnd, inordeIndexrMap),
		}
	}
	return helper(inorder, postorder, len(postorder)-1, 0, len(inorder)-1, inordeIndexrMap)
}

func buildTree1(inorder []int, postorder []int) *TreeNode {
	inorderMap := make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		inorderMap[inorder[i]] = i
	}
	return helper(inorder, postorder, 0, len(inorder)-1, len(postorder)-1, inorderMap)
}

func helper(inorder []int, postorder []int, inStart, inEnd, postIndex int, inorderMap map[int]int) *TreeNode {
	if postIndex < 0 || inStart > inEnd {
		return nil
	}
	root := &TreeNode{Val: postorder[postIndex]}
	rootIndex := inorderMap[postorder[postIndex]]
	rightNodeTotal := inEnd - rootIndex
	root.Left = helper(inorder, postorder, inStart, rootIndex-1, postIndex-rightNodeTotal-1, inorderMap)
	root.Right = helper(inorder, postorder, rootIndex+1, inEnd, postIndex-1, inorderMap)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
