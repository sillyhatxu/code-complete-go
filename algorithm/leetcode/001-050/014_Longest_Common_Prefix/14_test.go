package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_function(t *testing.T) {
	fmt.Println(strings.HasPrefix("flow", "flower"))
	fmt.Println(strings.HasPrefix("flight", "flower"))
	fmt.Println(strings.HasPrefix("flight", "fl"))
}
func Test_longestCommonPrefix(t *testing.T) {
	assert.EqualValues(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	assert.EqualValues(t, "", longestCommonPrefix([]string{"dog", "racecar", "car"}))
	assert.EqualValues(t, "", longestCommonPrefix([]string{}))
	assert.EqualValues(t, "a", longestCommonPrefix([]string{"aa", "a"}))
}

func Test_longestCommonPrefixOriginal(t *testing.T) {
	assert.EqualValues(t, "fl", longestCommonPrefixOriginal([]string{"flower", "flow", "flight"}))
	assert.EqualValues(t, "", longestCommonPrefixOriginal([]string{"dog", "racecar", "car"}))
	assert.EqualValues(t, "", longestCommonPrefixOriginal([]string{}))
	assert.EqualValues(t, "a", longestCommonPrefixOriginal([]string{"aa", "a"}))
}
