package _42_Linked_List_Cycle_II

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	check := head
	for slow != check {
		slow = slow.Next
		check = check.Next
	}
	return slow
}

type ListNode struct {
	Val  int
	Next *ListNode
}
