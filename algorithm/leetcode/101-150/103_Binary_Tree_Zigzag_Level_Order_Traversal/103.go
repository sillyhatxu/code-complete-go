package _03_Binary_Tree_Zigzag_Level_Order_Traversal

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var leftStack []*TreeNode
	var rightStack []*TreeNode
	leftStack = append(leftStack, root)
	for len(leftStack) > 0 || len(rightStack) > 0 {
		var currentRes []int
		for len(leftStack) > 0 {
			currentNode := leftStack[0]
			leftStack = leftStack[1:]
			currentRes = append(currentRes, currentNode.Val)
			if currentNode.Left != nil {
				rightStack = append([]*TreeNode{currentNode.Left}, rightStack...)
			}
			if currentNode.Right != nil {
				rightStack = append([]*TreeNode{currentNode.Right}, rightStack...)
			}
		}
		if len(currentRes) > 0 {
			result = append(result, currentRes)
		}
		currentRes = []int{}
		for len(rightStack) > 0 {
			currentNode := rightStack[0]
			rightStack = rightStack[1:]
			currentRes = append(currentRes, currentNode.Val)
			if currentNode.Right != nil {
				leftStack = append([]*TreeNode{currentNode.Right}, leftStack...)
			}
			if currentNode.Left != nil {
				leftStack = append([]*TreeNode{currentNode.Left}, leftStack...)
			}
		}
		if len(currentRes) > 0 {
			result = append(result, currentRes)
		}
	}
	return result
}

func zigzagLevelOrder1(root *TreeNode) [][]int {
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
