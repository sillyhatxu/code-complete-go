package _64_Minimum_Path_Sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minPathSum(t *testing.T) {
	assert.EqualValues(t, 7, minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
