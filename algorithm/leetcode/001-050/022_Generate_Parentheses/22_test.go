package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	assert.EqualValues(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParenthesis(3))
}

func Test_generateParenthesisOriginal(t *testing.T) {
	assert.EqualValues(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParenthesisOriginal(3))
}
