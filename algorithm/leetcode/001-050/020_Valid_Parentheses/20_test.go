package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid(t *testing.T) {
	assert.EqualValues(t, true, isValid("()"))
	assert.EqualValues(t, true, isValid("()[]{}"))
	assert.EqualValues(t, false, isValid("(]"))
	assert.EqualValues(t, false, isValid("([)]"))
	assert.EqualValues(t, true, isValid("{[]}"))
}

func Test_isValidOriginal(t *testing.T) {
	assert.EqualValues(t, true, isValidOriginal("()"))
	assert.EqualValues(t, true, isValidOriginal("()[]{}"))
	assert.EqualValues(t, false, isValidOriginal("(]"))
	assert.EqualValues(t, false, isValidOriginal("([)]"))
	assert.EqualValues(t, true, isValidOriginal("{[]}"))
}
