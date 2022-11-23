package _06_Reverse_Linked_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseListRecursion(head.Next)
	nextHead := head.Next
	nextHead.Next = head
	head.Next = nil
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}
