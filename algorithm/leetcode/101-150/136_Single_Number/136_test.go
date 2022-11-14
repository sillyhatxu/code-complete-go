package _36_Single_Number

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber1(t *testing.T) {
	fmt.Println(3 ^ 4)
}

func Test_singleNumber(t *testing.T) {
	assert.EqualValues(t, 2, singleNumber([]int{2, 2, 1}))
	assert.EqualValues(t, 4, singleNumber([]int{4, 1, 2, 1, 2}))
}
