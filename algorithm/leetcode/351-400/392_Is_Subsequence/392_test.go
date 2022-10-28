package _92_Is_Subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	expected1 := isSubsequence("abc", "ahbgdc")
	assert.EqualValues(t, true, expected1)
	expected2 := isSubsequence("axc", "ahbgdc")
	assert.EqualValues(t, false, expected2)

}
