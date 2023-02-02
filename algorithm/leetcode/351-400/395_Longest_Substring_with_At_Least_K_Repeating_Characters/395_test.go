package _95_Longest_Substring_with_At_Least_K_Repeating_Characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestSubstring(t *testing.T) {
	assert.EqualValues(t, 6, longestSubstring("aaabbb", 3))
	assert.EqualValues(t, 6, longestSubstring("aaabbbcaabb", 2))
}

func Test_longestSubstring1(t *testing.T) {
	assert.EqualValues(t, 6, longestSubstring1("aaabbbcaabb", 2))
}
