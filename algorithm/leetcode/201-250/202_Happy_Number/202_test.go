package _02_Happy_Number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isHappy(t *testing.T) {
	assert.EqualValues(t, true, isHappy(19))
}
