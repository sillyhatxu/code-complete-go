package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_strStr(t *testing.T) {
	assert.EqualValues(t, 0, strStr("sadbutsad", "sad"))
	assert.EqualValues(t, 0, strStr("a", "a"))
	assert.EqualValues(t, 2, strStr("hello", "ll"))
	assert.EqualValues(t, -1, strStr("aaaaa", "bba"))
}
