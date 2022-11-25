package _34_Palindrome_Linked_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	temp := &ListNode{Val: -1, Next: head}
	//          []    1    2    3    2    1
	//from here ↑
	fast := temp
	slow := temp
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//1    2    3    2    1
	//          ↑ if lise length is odd number. slow is here
	//1    2    2    1
	//     ↑ if lise length is even number. slow is here
	newNodeList := reverse(slow.Next)
	for head != nil && newNodeList != nil {
		if head.Val != newNodeList.Val {
			return false
		}
		head = head.Next
		newNodeList = newNodeList.Next
	}
	return true
}

func reverse(node *ListNode) *ListNode {
	var prev *ListNode
	for node != nil {
		node.Next, prev, node = prev, node, node.Next
	}
	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}
