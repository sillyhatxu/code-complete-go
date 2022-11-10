package _03_Binary_Tree_Zigzag_Level_Order_Traversal

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	stack1 := []*TreeNode{root}
	stack2 := []*TreeNode{}
	for len(stack1) > 0 || len(stack2) > 0 {
		var level []int
		for len(stack1) > 0 {
			current := stack1[0]
			stack1 = stack1[1:]
			level = append(level, current.Val)
			if current.Left != nil {
				stack2 = append([]*TreeNode{current.Left}, stack2...)
			}
			if current.Right != nil {
				stack2 = append([]*TreeNode{current.Right}, stack2...)
			}
		}
		result = append(result, level)
		level = []int{}
		for len(stack2) > 0 {
			current := stack2[0]
			stack2 = stack2[1:]
			level = append(level, current.Val)
			if current.Right != nil {
				stack1 = append([]*TreeNode{current.Right}, stack1...)
			}
			if current.Left != nil {
				stack1 = append([]*TreeNode{current.Left}, stack1...)
			}
		}
		if len(level) > 0 {
			result = append(result, level)
		}
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
