package _143_Longest_Common_Subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonSubsequence(t *testing.T) {
	assert.EqualValues(t, 3, longestCommonSubsequence("abc", "abc"))
	assert.EqualValues(t, 2, longestCommonSubsequence("ezupkr", "ubmrapg"))
	assert.EqualValues(t, 1, longestCommonSubsequence("bl", "yby"))
	assert.EqualValues(t, 3, longestCommonSubsequence("abcde", "ace"))
	assert.EqualValues(t, 0, longestCommonSubsequence("abc", "def"))
}
