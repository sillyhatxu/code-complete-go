package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	assert.EqualValues(t, -2147483647, myAtoi("-2147483647"))
	assert.EqualValues(t, -2147483648, myAtoi("-91283472332"))
	assert.EqualValues(t, 2147483647, myAtoi("9223372036854775808"))
	assert.EqualValues(t, 0, myAtoi("+-12"))
	assert.EqualValues(t, 0, myAtoi("words and 987"))
	assert.EqualValues(t, 42, myAtoi("42"))
	assert.EqualValues(t, -42, myAtoi("   -42"))
	assert.EqualValues(t, 4193, myAtoi("4193 with words"))
}
