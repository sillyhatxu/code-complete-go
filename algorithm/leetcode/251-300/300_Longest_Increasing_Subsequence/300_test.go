package _00_Longest_Increasing_Subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	assert.EqualValues(t, 4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
