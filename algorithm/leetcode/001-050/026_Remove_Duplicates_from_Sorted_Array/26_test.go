package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	assert.EqualValues(t, 2, removeDuplicates([]int{1, 1, 2}))
	assert.EqualValues(t, 5, removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func Test_removeDuplicates1(t *testing.T) {
	assert.EqualValues(t, 2, removeDuplicates1([]int{1, 1, 2}))
	assert.EqualValues(t, 5, removeDuplicates1([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func Test_removeDuplicates2(t *testing.T) {
	assert.EqualValues(t, 2, removeDuplicates2([]int{1, 1, 2}))
	assert.EqualValues(t, 5, removeDuplicates2([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
