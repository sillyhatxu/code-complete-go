package _28_Odd_Even_Linked_List

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	input := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	res := oddEvenList(input)
	assert.EqualValues(t, 1, res.Val)
	assert.EqualValues(t, 3, res.Next.Val)
	assert.EqualValues(t, 5, res.Next.Next.Val)
	assert.EqualValues(t, 2, res.Next.Next.Next.Val)
	assert.EqualValues(t, 4, res.Next.Next.Next.Next.Val)
}

func Test_oddEvenList1(t *testing.T) {
	input := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	res := oddEvenList1(input)
	assert.EqualValues(t, 1, res.Val)
	assert.EqualValues(t, 3, res.Next.Val)
	assert.EqualValues(t, 5, res.Next.Next.Val)
	assert.EqualValues(t, 2, res.Next.Next.Next.Val)
	assert.EqualValues(t, 4, res.Next.Next.Next.Next.Val)
}
