package _99_Binary_Tree_Right_Side_View

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		max := len(queue)
		for i := 0; i < max; i++ {
			curr := queue[i]
			if i == 0 {
				res = append(res, curr.Val)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
		}
		queue = queue[max:]
	}
	return res
}

func rightSideView1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		max := len(queue)
		for i := 0; i < max; i++ {
			curr := queue[0]
			queue = queue[1:]
			if i == 0 {
				res = append(res, curr.Val)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
		}
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
