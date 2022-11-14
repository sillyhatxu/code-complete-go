package _38_Copy_List_with_Random_Pointer

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*Node]*Node)
	currentNode := head
	for currentNode != nil {
		nodeMap[currentNode] = &Node{Val: currentNode.Val}
		currentNode = currentNode.Next
	}
	currentNode = head
	for currentNode != nil {
		nodeMap[currentNode].Next = nodeMap[currentNode.Next]
		nodeMap[currentNode].Random = nodeMap[currentNode.Random]
		currentNode = currentNode.Next
	}
	return nodeMap[head]
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
