package _71_Sum_of_Two_Integers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getSum(t *testing.T) {
	assert.EqualValues(t, 13, getSum(6, 7))
	assert.EqualValues(t, 50, getSum(20, 30))
}
