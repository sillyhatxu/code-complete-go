package _54_4Sum_II

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fourSumCount(t *testing.T) {
	assert.EqualValues(t, 6, fourSumCount([]int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1}))
	assert.EqualValues(t, 2, fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}
