package _50_Intersection_of_Two_Arrays_II

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_intersect(t *testing.T) {
	assert.EqualValues(t, []int{2, 2}, intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}
