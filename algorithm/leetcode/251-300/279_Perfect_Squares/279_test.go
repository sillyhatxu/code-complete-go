package _79_Perfect_Squares

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numSquares(t *testing.T) {
	assert.EqualValues(t, 3, numSquares(12))
}
