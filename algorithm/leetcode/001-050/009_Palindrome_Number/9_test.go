package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	assert.EqualValues(t, false, isPalindrome(123))
	assert.EqualValues(t, true, isPalindrome(121))
	assert.EqualValues(t, false, isPalindrome(-121))
	assert.EqualValues(t, false, isPalindrome(10))
}

func Test_isPalindromeHistory(t *testing.T) {
	assert.EqualValues(t, false, isPalindromeHistory(123))
	assert.EqualValues(t, true, isPalindromeHistory(121))
	assert.EqualValues(t, false, isPalindromeHistory(-121))
	assert.EqualValues(t, false, isPalindromeHistory(10))
}
