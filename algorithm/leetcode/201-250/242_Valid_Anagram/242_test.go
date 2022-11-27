package _42_Valid_Anagram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	assert.EqualValues(t, true, isAnagram("anagram", "nagaram"))
	assert.EqualValues(t, false, isAnagram("rat", "car"))
	assert.EqualValues(t, false, isAnagram("a", "ab"))
}
