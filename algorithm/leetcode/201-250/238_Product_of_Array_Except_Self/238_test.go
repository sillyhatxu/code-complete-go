package _38_Product_of_Array_Except_Self

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	assert.EqualValues(t, []int{}, productExceptSelf([]int{1, 2, 3, 4}))
}
