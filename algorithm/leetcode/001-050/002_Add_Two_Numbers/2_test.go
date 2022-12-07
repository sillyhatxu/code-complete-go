package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	result1 := addTwoNumbers(
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
		&ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	)
	assert.EqualValues(t, 7, result1.Val)
	assert.EqualValues(t, 0, result1.Next.Val)
	assert.EqualValues(t, 8, result1.Next.Next.Val)
	result2 := addTwoNumbers(
		&ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
								},
							},
						},
					},
				},
			},
		},
		&ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		},
	)
	assert.EqualValues(t, 8, result2.Val)
	assert.EqualValues(t, 9, result2.Next.Val)
	assert.EqualValues(t, 9, result2.Next.Next.Val)
	assert.EqualValues(t, 9, result2.Next.Next.Next.Val)
	assert.EqualValues(t, 0, result2.Next.Next.Next.Next.Val)
	assert.EqualValues(t, 0, result2.Next.Next.Next.Next.Next.Val)
	assert.EqualValues(t, 0, result2.Next.Next.Next.Next.Next.Next.Val)
	assert.EqualValues(t, 1, result2.Next.Next.Next.Next.Next.Next.Next.Val)
}
