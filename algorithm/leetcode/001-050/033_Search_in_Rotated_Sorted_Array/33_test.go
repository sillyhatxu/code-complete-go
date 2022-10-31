package _33_Search_in_Rotated_Sorted_Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_search(t *testing.T) {
	result1 := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	assert.EqualValues(t, 4, result1)
	result2 := search([]int{1}, 0)
	assert.EqualValues(t, -1, result2)
	result3 := search([]int{3, 1}, 1)
	assert.EqualValues(t, 1, result3)
	result4 := search([]int{3, 1}, 3)
	assert.EqualValues(t, 0, result4)
	result5 := search([]int{3, 4, 5, 6, 1, 2}, 2)
	assert.EqualValues(t, 5, result5)
	result6 := search([]int{1, 3}, 0)
	assert.EqualValues(t, -1, result6)
}
