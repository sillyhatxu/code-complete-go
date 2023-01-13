package _46_Permutations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_permute(t *testing.T) {
	assert.EqualValues(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, permute([]int{1, 2, 3}))
	assert.EqualValues(t, [][]int{{0, 1}, {1, 0}}, permute([]int{0, 1}))
}

func Test_permute1(t *testing.T) {
	assert.EqualValues(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, permute1([]int{1, 2, 3}))
	assert.EqualValues(t, [][]int{{0, 1}, {1, 0}}, permute1([]int{0, 1}))
}
