package _200_Minimum_Absolute_Difference

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumAbsDifference(t *testing.T) {
	expected1 := minimumAbsDifference([]int{4, 2, 1, 3})
	assert.EqualValues(t, expected1, [][]int{{1, 2}, {2, 3}, {3, 4}})

	expected2 := minimumAbsDifference([]int{1, 3, 6, 10, 15})
	assert.EqualValues(t, expected2, [][]int{{1, 3}})
}
