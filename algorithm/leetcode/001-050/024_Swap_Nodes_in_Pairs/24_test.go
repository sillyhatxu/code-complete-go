package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	input := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	swapPairs(input)
	assert.EqualValues(t, 2, input.Val)
	assert.EqualValues(t, 1, input.Next.Val)
	assert.EqualValues(t, 4, input.Next.Next.Val)
	assert.EqualValues(t, 3, input.Next.Next.Next.Val)
}
