package _45_Jump_Game_II

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_jump(t *testing.T) {
	assert.EqualValues(t, 2, jump([]int{2, 3, 1, 1, 4}))
	assert.EqualValues(t, 2, jump([]int{1, 2, 3}))
}
