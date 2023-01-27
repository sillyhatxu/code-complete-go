package _52_Maximum_Product_Subarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProduct(t *testing.T) {
	assert.EqualValues(t, 6, maxProduct([]int{2, 3, -2, 4}))
}

func Test_maxProduct1(t *testing.T) {
	assert.EqualValues(t, 6, maxProduct1([]int{2, 3, -2, 4}))
}
