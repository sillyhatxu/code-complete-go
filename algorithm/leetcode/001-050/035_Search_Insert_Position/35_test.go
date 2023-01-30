package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	assert.EqualValues(t, 2, searchInsert([]int{1, 3, 5, 6}, 5))
	assert.EqualValues(t, 1, searchInsert([]int{1, 3, 5, 6}, 2))
	assert.EqualValues(t, 4, searchInsert([]int{1, 3, 5, 6}, 7))
	assert.EqualValues(t, 0, searchInsert([]int{1, 3, 5, 6}, 0))
}

func Test_searchInsert1(t *testing.T) {
	assert.EqualValues(t, 2, searchInsert1([]int{1, 3, 5, 6}, 5))
	assert.EqualValues(t, 1, searchInsert1([]int{1, 3, 5, 6}, 2))
	assert.EqualValues(t, 4, searchInsert1([]int{1, 3, 5, 6}, 7))
	assert.EqualValues(t, 0, searchInsert1([]int{1, 3, 5, 6}, 0))
}
