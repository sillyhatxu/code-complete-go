package _40_Search_a_2D_Matrix_II

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	assert.EqualValues(t, true, searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	assert.EqualValues(t, false, searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
}
