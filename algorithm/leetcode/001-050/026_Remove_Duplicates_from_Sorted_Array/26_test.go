package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	assert.EqualValues(t, 2, removeDuplicates([]int{1, 1, 2}))
	assert.EqualValues(t, 5, removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
