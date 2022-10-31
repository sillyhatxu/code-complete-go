package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_nextPermutationOptimize(t *testing.T) {
	debug := []int{1, 3, 2}
	nextPermutationOptimize(debug)
	assert.EqualValues(t, []int{2, 1, 3}, debug)
	input := []int{3, 2, 1}
	nextPermutationOptimize(input)
	assert.EqualValues(t, []int{1, 2, 3}, input)
	input1 := []int{1, 2, 3}
	nextPermutationOptimize(input1)
	assert.EqualValues(t, []int{1, 3, 2}, input1)
	input2 := []int{3, 2, 1}
	nextPermutationOptimize(input2)
	assert.EqualValues(t, []int{1, 2, 3}, input2)
	input3 := []int{1, 1, 5}
	nextPermutationOptimize(input3)
	assert.EqualValues(t, []int{1, 5, 1}, input3)
	input4 := []int{1, 2, 3, 9, 5, 5, 2, 1}
	nextPermutationOptimize(input4)
	assert.EqualValues(t, []int{1, 2, 5, 1, 2, 3, 5, 9}, input4)
}
