package _30_Surrounded_Regions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solve(t *testing.T) {
	input1 := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	expected1 := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(input1)
	assert.EqualValues(t, expected1, input1)
}
