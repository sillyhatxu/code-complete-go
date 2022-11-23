package _00_Number_of_Islands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numIslands(t *testing.T) {

	assert.EqualValues(t, 1, numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
}
