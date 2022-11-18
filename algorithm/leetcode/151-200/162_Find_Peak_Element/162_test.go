package _62_Find_Peak_Element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findPeakElement(t *testing.T) {
	assert.EqualValues(t, 2, findPeakElement([]int{1, 2, 3, 1}))
	assert.EqualValues(t, 2, findPeakElement([]int{1, 2, 3, 1}))
	assert.EqualValues(t, 1, findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	assert.EqualValues(t, 1, findPeakElement([]int{3, 2, 1}))
}
