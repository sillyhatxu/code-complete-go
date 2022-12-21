package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_intToRoman(t *testing.T) {
	assert.EqualValues(t, "LXXX", intToRoman(80))
	assert.EqualValues(t, "CD", intToRoman(400))
	assert.EqualValues(t, "IV", intToRoman(4))
	assert.EqualValues(t, "IX", intToRoman(9))
	assert.EqualValues(t, "XL", intToRoman(40))
	assert.EqualValues(t, "LVIII", intToRoman(58))
	assert.EqualValues(t, "MCMXCIV", intToRoman(1994))
}

func Test_intToRomanOptimize(t *testing.T) {
	assert.EqualValues(t, "CD", intToRomanOptimize(400))
	assert.EqualValues(t, "IV", intToRomanOptimize(4))
	assert.EqualValues(t, "IX", intToRomanOptimize(9))
	assert.EqualValues(t, "XL", intToRomanOptimize(40))
	assert.EqualValues(t, "LVIII", intToRomanOptimize(58))
	assert.EqualValues(t, "MCMXCIV", intToRomanOptimize(1994))
}
