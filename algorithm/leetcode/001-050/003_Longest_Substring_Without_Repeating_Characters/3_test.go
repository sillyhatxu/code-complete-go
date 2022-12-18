package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	assert.EqualValues(t, 2, lengthOfLongestSubstring("abba"))
	assert.EqualValues(t, 1, lengthOfLongestSubstring(" "))
	assert.EqualValues(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.EqualValues(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.EqualValues(t, 3, lengthOfLongestSubstring("pwwkew"))
}
