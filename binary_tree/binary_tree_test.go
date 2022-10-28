package binary_tree

import (
	"fmt"
	"testing"
)

/**
深度优先遍历 - 前序遍历：
F, B, A, D, C, E, G, I, H
*/
func TestPreOrderTraversal(t *testing.T) {
	demo := getTree()
	fmt.Println("pre-order traversal")
	preOrderTraversal(demo)
	fmt.Println()
}

/**
深度优先遍历 - 中序遍历：
A, B, C, D, E, F, G, H, I.
*/
func TestInOrderTraversal(t *testing.T) {
	demo := getTree()
	fmt.Println("in-order traversal")
	inOrderTraversal(demo)
	fmt.Println()
}

/**
深度优先搜索 - 后序遍历：
A, C, E, D, B, H, I, G, F.
*/
func TestPostOrderTraversal(t *testing.T) {
	demo := getTree()
	fmt.Println("post-order traversal")
	postOrderTraversal(demo)
	fmt.Println()
}

/**
广度优先遍历 - 层次遍历：
F, B, G, A, D, I, C, E, H.
*/
func TestBreadthFirstTraversal(t *testing.T) {
	demo := getTree()
	fmt.Println("breadth-first traversal")
	breadthFirstTraversal(demo)
	fmt.Println()
}
