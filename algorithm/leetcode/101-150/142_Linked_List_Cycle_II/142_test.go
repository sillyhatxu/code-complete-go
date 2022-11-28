package _42_Linked_List_Cycle_II

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	node0 := &ListNode{Val: 3}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 0}
	node3 := &ListNode{Val: -4}
	node0.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node1
	assert.EqualValues(t, node1, detectCycle(node0))

	node00 := &ListNode{Val: 3}
	node01 := &ListNode{Val: 2}
	node00.Next = node01
	node01.Next = node00
	assert.EqualValues(t, node00, detectCycle(node00))

	node11 := &ListNode{Val: 3}
	assert.Nil(t, detectCycle(node11))
}
