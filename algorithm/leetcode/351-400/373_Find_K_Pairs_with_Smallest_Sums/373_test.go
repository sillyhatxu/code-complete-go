package _73_Find_K_Pairs_with_Smallest_Sums

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_kSmallestPairs(t *testing.T) {
	assert.EqualValues(t, [][]int{{1, 2}, {1, 4}, {1, 6}}, kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
}
