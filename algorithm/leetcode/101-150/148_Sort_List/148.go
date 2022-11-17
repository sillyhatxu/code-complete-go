package _48_Sort_List

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMid(head)
	midNext := mid.Next
	mid.Next = nil
	return merge(sortList(head), sortList(midNext))
}

func getMid(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func merge(left *ListNode, right *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	for left != nil && right != nil {
		if left.Val > right.Val {
			current.Next = right
			right = right.Next
		} else {
			current.Next = left
			left = left.Next
		}
		current = current.Next
	}
	if left != nil {
		current.Next = left
	} else if right != nil {
		current.Next = right
	}
	return result.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
