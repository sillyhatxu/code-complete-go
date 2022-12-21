package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	assert.EqualValues(t, 3, romanToInt("III"))
	assert.EqualValues(t, 4, romanToInt("IV"))
	assert.EqualValues(t, 9, romanToInt("IX"))
	assert.EqualValues(t, 58, romanToInt("LVIII"))
	assert.EqualValues(t, 1994, romanToInt("MCMXCIV"))
}

func Test_romanToIntOriginal(t *testing.T) {
	assert.EqualValues(t, 3, romanToIntOriginal("III"))
	assert.EqualValues(t, 4, romanToIntOriginal("IV"))
	assert.EqualValues(t, 9, romanToIntOriginal("IX"))
	assert.EqualValues(t, 58, romanToIntOriginal("LVIII"))
	assert.EqualValues(t, 1994, romanToIntOriginal("MCMXCIV"))
}
