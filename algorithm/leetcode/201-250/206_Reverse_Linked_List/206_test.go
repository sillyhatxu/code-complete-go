package _06_Reverse_Linked_List

import (
	"fmt"
	"testing"
)

func Test_reverseList_iteration(t *testing.T) {
	res := reverseList(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	})
	iterator(res)
}

func Test_reverseList_recursion(t *testing.T) {
	res := reverseListRecursion(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	})
	iterator(res)
}

func iterator(node *ListNode) {
	if node != nil {
		fmt.Print(node.Val, "->")
		iterator(node.Next)
	}
	fmt.Println()
}
