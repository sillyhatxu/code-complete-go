package _39_Combination_Sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	assert.EqualValues(t, [][]int{{2, 2, 3}, {7}}, combinationSum([]int{2, 3, 6, 7}, 7))
}
