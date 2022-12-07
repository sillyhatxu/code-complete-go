package _15_3Sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_threeSum(t *testing.T) {
	assert.EqualValues(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
