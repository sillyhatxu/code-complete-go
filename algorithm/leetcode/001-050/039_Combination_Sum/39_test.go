package _39_Combination_Sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	assert.EqualValues(t, [][]int{{2, 2, 3}, {7}}, combinationSum([]int{2, 3, 6, 7}, 7))
	assert.EqualValues(t, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, combinationSum([]int{2, 3, 5}, 8))
}

func Test_combinationSum1(t *testing.T) {
	assert.EqualValues(t, [][]int{{2, 2, 3}, {7}}, combinationSum1([]int{2, 3, 6, 7}, 7))
	assert.EqualValues(t, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, combinationSum1([]int{2, 3, 5}, 8))
}

func Test_combinationSum2(t *testing.T) {
	assert.EqualValues(t, [][]int{{2, 2, 3}, {7}}, combinationSum2([]int{2, 3, 6, 7}, 7))
	assert.EqualValues(t, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, combinationSum2([]int{2, 3, 5}, 8))
}
