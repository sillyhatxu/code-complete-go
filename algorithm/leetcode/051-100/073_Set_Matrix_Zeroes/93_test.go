package _73_Set_Matrix_Zeroes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	input1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(input1)
	expected1 := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	assert.EqualValues(t, expected1, input1)

	input2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(input2)
	expected2 := [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}
	assert.EqualValues(t, expected2, input2)

	input3 := [][]int{{1, 0, 3}}
	setZeroes(input3)
	expected3 := [][]int{{0, 0, 0}}
	assert.EqualValues(t, expected3, input3)

	input4 := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 0}}
	setZeroes(input4)
	expected4 := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 0}}
	assert.EqualValues(t, expected4, input4)

	input5 := [][]int{{1, 1, 1, 1, 0}, {1, 2, 2, 2, 2}, {1, 2, 2, 2, 2}, {1, 2, 0, 2, 2}, {1, 2, 2, 0, 2}}
	setZeroes(input5)
	expected5 := [][]int{{0, 0, 0, 0, 0}, {1, 2, 0, 0, 0}, {1, 2, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	assert.EqualValues(t, expected5, input5)
}
