package _16_Populating_Next_Right_Pointers_in_Each_Node

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_connect(t *testing.T) {
	res := connect(&Node{
		Val:   1,
		Left:  &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}},
		Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}},
	})
	assert.EqualValues(t, 1, res.Val)
	assert.Nil(t, res.Next)

	assert.EqualValues(t, 2, res.Left.Val)
	assert.EqualValues(t, 3, res.Left.Next.Val)
	assert.Nil(t, res.Left.Next.Next)

	assert.EqualValues(t, 4, res.Left.Left.Val)
	assert.EqualValues(t, 5, res.Left.Left.Next.Val)
	assert.EqualValues(t, 6, res.Left.Left.Next.Next.Val)
	assert.EqualValues(t, 7, res.Left.Left.Next.Next.Next.Val)
	assert.Nil(t, res.Left.Left.Next.Next.Next.Next)
}

func Test_connect1(t *testing.T) {
	res := connect1(&Node{
		Val:   1,
		Left:  &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}},
		Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}},
	})
	assert.EqualValues(t, 1, res.Val)
	assert.Nil(t, res.Next)

	assert.EqualValues(t, 2, res.Left.Val)
	assert.EqualValues(t, 3, res.Left.Next.Val)
	assert.Nil(t, res.Left.Next.Next)

	assert.EqualValues(t, 4, res.Left.Left.Val)
	assert.EqualValues(t, 5, res.Left.Left.Next.Val)
	assert.EqualValues(t, 6, res.Left.Left.Next.Next.Val)
	assert.EqualValues(t, 7, res.Left.Left.Next.Next.Next.Val)
	assert.Nil(t, res.Left.Left.Next.Next.Next.Next)
}
