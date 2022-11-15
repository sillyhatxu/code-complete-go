package _41_Linked_List_Cycle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}
	node7 := &ListNode{Val: 7}
	node8 := &ListNode{Val: 8}
	node9 := &ListNode{Val: 9}
	test1 := func() {
		input := &ListNode{Val: 0, Next: node1}
		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node5
		node5.Next = node6
		node6.Next = node7
		node7.Next = node8
		node8.Next = node9
		node9.Next = node1
		assert.EqualValues(t, true, hasCycle(input))
	}
	test1()

	test2 := func() {
		input := &ListNode{Val: 0, Next: node1}
		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node5
		node5.Next = node6
		node6.Next = node7
		node7.Next = node8
		node8.Next = node9
		node9.Next = nil
		assert.EqualValues(t, false, hasCycle(input))
	}
	test2()

	test3 := func() {
		input := &ListNode{Val: 0, Next: node1}
		node1.Next = node2
		node2.Next = node3
		node3.Next = node1
		assert.EqualValues(t, true, hasCycle(input))
	}
	test3()
}
