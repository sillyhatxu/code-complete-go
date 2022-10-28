package binary_tree

import "fmt"

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

func getTree() *TreeNode {
	return &TreeNode{
		Val: "F",
		Left: &TreeNode{
			Val: "B",
			Left: &TreeNode{
				Val:   "A",
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: "D",
				Left: &TreeNode{
					Val: "C",
				},
				Right: &TreeNode{
					Val: "E",
				},
			},
		},
		Right: &TreeNode{
			Val: "G",
			Right: &TreeNode{
				Val: "I",
				Left: &TreeNode{
					Val: "H",
				},
			},
		},
	}
}

func preOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Val, ", ")
	preOrderTraversal(node.Left)
	preOrderTraversal(node.Right)
}

func inOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	inOrderTraversal(node.Left)
	fmt.Print(node.Val, ", ")
	inOrderTraversal(node.Right)
}

func postOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	postOrderTraversal(node.Left)
	postOrderTraversal(node.Right)
	fmt.Print(node.Val, ", ")
}

func breadthFirstTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	var result []*TreeNode
	result = append(result, node)

	i := 0
	for result != nil && len(result) > 0 {
		treeNode := result[i]
		result = append(result[:i], result[i+1:]...)
		fmt.Print(treeNode.Val, ", ")
		leftNode := treeNode.Left
		if leftNode != nil {
			result = append(result, leftNode)
		}
		rightNode := treeNode.Right
		if rightNode != nil {
			result = append(result, rightNode)
		}

	}
}
