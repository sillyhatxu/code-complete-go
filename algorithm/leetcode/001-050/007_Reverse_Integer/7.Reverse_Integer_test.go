package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverse(t *testing.T) {
	assert.EqualValues(t, 0, reverse(1534236469))
	assert.EqualValues(t, 321, reverse(123))
	assert.EqualValues(t, -321, reverse(-123))
	assert.EqualValues(t, 21, reverse(120))
}

func Test_reverse1(t *testing.T) {
	assert.EqualValues(t, 321, reverse1(123))
	assert.EqualValues(t, -321, reverse1(-123))
	assert.EqualValues(t, 21, reverse1(120))
}
