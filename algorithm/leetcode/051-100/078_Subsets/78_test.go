package _78_Subsets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subsets(t *testing.T) {
	input := []int{1, 2, 3}
	expected := [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}
	result := subsets(input)
	assert.EqualValues(t, expected, result)
}

func Test_subsets1(t *testing.T) {
	input := []int{1, 2, 3}
	expected := [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}
	result := subsets1(input)
	assert.EqualValues(t, expected, result)
}
