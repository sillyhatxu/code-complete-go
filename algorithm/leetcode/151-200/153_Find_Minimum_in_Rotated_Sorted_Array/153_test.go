package _53_Find_Minimum_in_Rotated_Sorted_Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMin(t *testing.T) {
	assert.EqualValues(t, 1, findMin([]int{3, 4, 5, 1, 2}))
	assert.EqualValues(t, 0, findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	assert.EqualValues(t, 11, findMin([]int{11, 13, 15, 17}))
}
