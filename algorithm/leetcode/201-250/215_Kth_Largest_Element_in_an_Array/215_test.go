package _15_Kth_Largest_Element_in_an_Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	assert.EqualValues(t, 5, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
