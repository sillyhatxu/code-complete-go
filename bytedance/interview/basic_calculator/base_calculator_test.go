package basic_calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	assert.EqualValues(t, 2147483647, calculate("2147483647"))
	assert.EqualValues(t, 7, calculate("1+2*3"))
	assert.EqualValues(t, 3, calculate(" 2-1 + 2 "))
	assert.EqualValues(t, 23, calculate("(1+(4+5+2)-3)+(6+8)"))
	assert.EqualValues(t, 6, calculate("1+2+3"))
}

func TestCalculate1(t *testing.T) {
	//assert.EqualValues(t, 2147483647, calculate1("2147483647"))
	assert.EqualValues(t, 7, calculate1("1+2*3"))
	assert.EqualValues(t, 33, calculate1("11+22"))
	assert.EqualValues(t, 3, calculate1(" 2-1 + 2 "))
	assert.EqualValues(t, 23, calculate1("(1+(4+5+2)-3)+(6+8)"))
	assert.EqualValues(t, 6, calculate1("1+2+3"))
}
