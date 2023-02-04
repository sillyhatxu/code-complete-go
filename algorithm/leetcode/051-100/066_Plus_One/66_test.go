package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_plusOne(t *testing.T) {
	assert.EqualValues(t, []int{1, 0, 0, 0}, plusOne([]int{9, 9, 9}))
	assert.EqualValues(t, []int{1, 0}, plusOne([]int{9}))
	assert.EqualValues(t, []int{1, 2, 4}, plusOne([]int{1, 2, 3}))
	assert.EqualValues(t, []int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))
}

func Test_plusOne1(t *testing.T) {
	assert.EqualValues(t, []int{1, 0, 0, 0}, plusOne1([]int{9, 9, 9}))
	assert.EqualValues(t, []int{1, 0}, plusOne1([]int{9}))
	assert.EqualValues(t, []int{1, 2, 4}, plusOne1([]int{1, 2, 3}))
	assert.EqualValues(t, []int{4, 3, 2, 2}, plusOne1([]int{4, 3, 2, 1}))
}

func Test_plusOne2(t *testing.T) {
	assert.EqualValues(t, []int{1, 0, 0, 0}, plusOne2([]int{9, 9, 9}))
	assert.EqualValues(t, []int{1, 0}, plusOne2([]int{9}))
	assert.EqualValues(t, []int{1, 2, 4}, plusOne2([]int{1, 2, 3}))
	assert.EqualValues(t, []int{4, 3, 2, 2}, plusOne2([]int{4, 3, 2, 1}))
}
