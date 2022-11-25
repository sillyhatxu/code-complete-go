package _34_Palindrome_Linked_List

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	assert.EqualValues(t, true, isPalindrome(&ListNode{Val: 1, Next: nil}))
	assert.EqualValues(t, true, isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}))
	assert.EqualValues(t, true, isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}))
}
