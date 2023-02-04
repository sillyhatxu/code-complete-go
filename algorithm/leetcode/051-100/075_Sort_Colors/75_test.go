package _75_Sort_Colors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortColors(t *testing.T) {
	input1 := []int{2, 0, 2, 1, 1, 0}
	sortColors(input1)
	assert.EqualValues(t, []int{0, 0, 1, 1, 2, 2}, input1)
	input2 := []int{2, 0, 1}
	sortColors(input2)
	assert.EqualValues(t, []int{0, 1, 2}, input2)
	input3 := []int{1, 2, 0}
	sortColors(input3)
	assert.EqualValues(t, []int{0, 1, 2}, input3)
}

func Test_sortColors1(t *testing.T) {
	input1 := []int{2, 0, 2, 1, 1, 0}
	sortColors1(input1)
	assert.EqualValues(t, []int{0, 0, 1, 1, 2, 2}, input1)
	input2 := []int{2, 0, 1}
	sortColors1(input2)
	assert.EqualValues(t, []int{0, 1, 2}, input2)
	input3 := []int{1, 2, 0}
	sortColors1(input3)
	assert.EqualValues(t, []int{0, 1, 2}, input3)
}
