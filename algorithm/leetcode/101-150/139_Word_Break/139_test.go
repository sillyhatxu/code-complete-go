package _39_Word_Break

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	assert.EqualValues(t, true, wordBreak("leetcode", []string{"leet", "code"}))
}
