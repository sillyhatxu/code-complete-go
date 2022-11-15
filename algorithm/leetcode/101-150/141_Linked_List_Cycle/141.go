package _41_Linked_List_Cycle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	rabbit, tortoise := head, head
	for rabbit.Next != nil && rabbit.Next.Next != nil {
		tortoise = tortoise.Next
		rabbit = rabbit.Next.Next
		if rabbit == tortoise {
			return true
		}
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}
