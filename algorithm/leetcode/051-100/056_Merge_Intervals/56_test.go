package _56_Merge_Intervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_merge(t *testing.T) {
	assert.EqualValues(t, [][]int{{1, 6}, {8, 10}, {15, 18}}, merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	assert.EqualValues(t, [][]int{{1, 5}}, merge([][]int{{1, 4}, {4, 5}}))
}
