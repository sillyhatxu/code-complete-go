package _74_Search_a_2D_Matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	assert.EqualValues(t, false, searchMatrix([][]int{{1, 1}}, 0))
	assert.EqualValues(t, true, searchMatrix([][]int{{1, 1}}, 1))
	assert.EqualValues(t, true, searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	assert.EqualValues(t, false, searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
}
