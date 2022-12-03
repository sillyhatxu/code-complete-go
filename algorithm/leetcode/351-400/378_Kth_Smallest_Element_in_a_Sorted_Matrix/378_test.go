package _78_Kth_Smallest_Element_in_a_Sorted_Matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	assert.EqualValues(t, 13, kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8))
}
