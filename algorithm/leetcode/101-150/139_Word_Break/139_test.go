package _39_Word_Break

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	assert.EqualValues(t, true, wordBreak("aleetcode", []string{"leet", "code"}))
	assert.EqualValues(t, true, wordBreak("leetcode", []string{"leet", "code"}))
	assert.EqualValues(t, true, wordBreak("applepenapple", []string{"apple", "pen"}))
	assert.EqualValues(t, false, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	assert.EqualValues(t, true, wordBreak("catanddogsandcatsand", []string{"cats", "dog", "sand", "and", "cat"}))
}
