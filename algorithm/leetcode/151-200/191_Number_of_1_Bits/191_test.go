package _91_Number_of_1_Bits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	assert.EqualValues(t, 3, hammingWeight(uint32(11)))
}
