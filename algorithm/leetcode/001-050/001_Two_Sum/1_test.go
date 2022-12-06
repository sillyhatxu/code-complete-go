package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum(t *testing.T) {
	assert.EqualValues(t, []int{1, 2}, twoSum([]int{2, 7, 11, 15}, 18))
}
