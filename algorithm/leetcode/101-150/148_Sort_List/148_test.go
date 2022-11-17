package _48_Sort_List

import (
	"fmt"
	"testing"
)

func Test_sortList(t *testing.T) {
	result1 := sortList(&ListNode{
		Val: -1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 0,
					},
				},
			},
		},
	})
	print(result1)
}

func print(list *ListNode) {
	for list != nil {
		fmt.Print(fmt.Sprintf("%d -> ", list.Val))
		list = list.Next
	}
	fmt.Println()
}
