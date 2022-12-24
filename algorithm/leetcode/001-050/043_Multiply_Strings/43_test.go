package _43_Multiply_Strings

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_multiply1(t *testing.T) {
	fmt.Println('0' - '0')
	fmt.Println('0' - '0')
	fmt.Println('1' - '0')
}

func Test_multiply(t *testing.T) {
	assert.EqualValues(t, "6", multiply("2", "3"))
	assert.EqualValues(t, "56088", multiply("123", "456"))
	assert.EqualValues(t, "998001", multiply("999", "999"))
}
