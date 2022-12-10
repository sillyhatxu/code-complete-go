package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeElement(t *testing.T) {
	assert.EqualValues(t, 2, removeElement([]int{3, 2, 2, 3}, 3))
	assert.EqualValues(t, 5, removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
