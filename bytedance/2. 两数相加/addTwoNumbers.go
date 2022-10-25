package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{} //返回结果
	current := &result
	carryBit := 0 //进位标示
	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
		}
		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
		}
		sum := v1 + v2 + carryBit
		current.Next = &ListNode{sum % 10, nil}
		carryBit = sum / 10
		current = current.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carryBit != 0 {
		current.Next = &ListNode{1, nil}
	}
	return result.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
