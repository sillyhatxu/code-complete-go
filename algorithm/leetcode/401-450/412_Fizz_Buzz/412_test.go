package _12_Fizz_Buzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	assert.EqualValues(t, []string{"1", "2", "Fizz"}, fizzBuzz(3))
}
