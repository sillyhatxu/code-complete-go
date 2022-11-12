package _25_Valid_Palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	assert.EqualValues(t, true, isPalindrome("A man, a plan, a canal: Panama"))
	assert.EqualValues(t, false, isPalindrome("race a car"))
}
