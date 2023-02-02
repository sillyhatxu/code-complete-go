package _5_Jump_Game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canJump(t *testing.T) {
	assert.EqualValues(t, true, canJump([]int{2, 3, 1, 1, 4}))
	assert.EqualValues(t, true, canJump([]int{2, 5, 0, 0}))
}

func Test_canJump1(t *testing.T) {
	assert.EqualValues(t, true, canJump1([]int{2, 3, 1, 1, 4}))
	assert.EqualValues(t, true, canJump1([]int{2, 5, 0, 0}))
}
