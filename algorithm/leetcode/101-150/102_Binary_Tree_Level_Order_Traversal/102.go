package _02_Binary_Tree_Level_Order_Traversal

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var currentLevel []int
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[i]
			currentLevel = append(currentLevel, curr.Val)
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		queue = queue[size:]
		res = append(res, currentLevel)
	}
	return res
}

func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var level []int
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			current := queue[0]
			queue = queue[1:]
			level = append(level, current.Val)
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
