package _26_Power_of_Three

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_isPowerOfThree1(t *testing.T) {
	temp := math.Log10(float64(27)) / math.Log10(float64(3))
	_, res := math.Modf(temp)
	fmt.Println(res)
}
func Test_isPowerOfThree(t *testing.T) {
	assert.EqualValues(t, true, isPowerOfThree(3))
	assert.EqualValues(t, true, isPowerOfThree(27))
	assert.EqualValues(t, false, isPowerOfThree(0))
	assert.EqualValues(t, true, isPowerOfThree(9))
	assert.EqualValues(t, false, isPowerOfThree(45))
}
