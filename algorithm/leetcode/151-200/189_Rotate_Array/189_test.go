package _89_Rotate_Array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_test(t *testing.T) {
	fmt.Println(3 % 7)
	fmt.Println(8 % 7)
}
func Test_rotate(t *testing.T) {
	input1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(input1, 3)
	assert.EqualValues(t, []int{5, 6, 7, 1, 2, 3, 4}, input1)

	input2 := []int{-1, -100, 3, 99}
	rotate(input2, 3)
	assert.EqualValues(t, []int{3, 99, -1, -100}, input2)
}
