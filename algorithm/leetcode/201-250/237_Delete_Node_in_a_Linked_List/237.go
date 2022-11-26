package _37_Delete_Node_in_a_Linked_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	node = node.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
