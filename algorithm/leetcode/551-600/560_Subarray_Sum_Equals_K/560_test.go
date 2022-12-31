package _60_Subarray_Sum_Equals_K

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subarraySum(t *testing.T) {
	assert.EqualValues(t, 5, subarraySum([]int{-2, 1, 1, 1, 2, 1, -2}, 3))
	assert.EqualValues(t, 4, subarraySum([]int{-2, 1, 3, 5, -1, -2, 6}, 4))
	assert.EqualValues(t, 2, subarraySum([]int{1, 1, 1}, 2))
	assert.EqualValues(t, 2, subarraySum([]int{1, 2, 3}, 3))
}
