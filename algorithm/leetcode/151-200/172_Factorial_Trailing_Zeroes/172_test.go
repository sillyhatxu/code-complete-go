package _72_Factorial_Trailing_Zeroes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	assert.EqualValues(t, 7, trailingZeroes(30))
	assert.EqualValues(t, 2, trailingZeroes(10))
	assert.EqualValues(t, 1, trailingZeroes(5))
	assert.EqualValues(t, 0, trailingZeroes(3))
	assert.EqualValues(t, 0, trailingZeroes(0))
}
