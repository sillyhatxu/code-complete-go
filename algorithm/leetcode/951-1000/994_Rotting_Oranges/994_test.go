package _94_Rotting_Oranges

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_orangesRotting(t *testing.T) {
	assert.EqualValues(t, 4, orangesRotting([][]int{
		{2, 1, 1}, {1, 1, 0}, {0, 1, 1},
	}))
	assert.EqualValues(t, -1, orangesRotting([][]int{
		{2, 1, 1}, {0, 1, 1}, {1, 0, 1},
	}))
	assert.EqualValues(t, 0, orangesRotting([][]int{
		{0, 2},
	}))
}
