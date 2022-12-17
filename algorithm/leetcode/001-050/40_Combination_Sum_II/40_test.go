package _0_Combination_Sum_II

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	assert.EqualValues(t, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	assert.EqualValues(t, [][]int{{1, 2, 2}, {5}}, combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
