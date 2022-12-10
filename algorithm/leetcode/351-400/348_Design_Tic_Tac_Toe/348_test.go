package _48_Design_Tic_Tac_Toe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TicTacToe_3(t *testing.T) {
	ticTacToe := Constructor(3)
	assert.EqualValues(t, 0, ticTacToe.Move(0, 0, 1))
	assert.EqualValues(t, 0, ticTacToe.Move(0, 2, 2))
	assert.EqualValues(t, 0, ticTacToe.Move(2, 2, 1))
	assert.EqualValues(t, 0, ticTacToe.Move(1, 1, 2))
	assert.EqualValues(t, 0, ticTacToe.Move(2, 0, 1))
	assert.EqualValues(t, 0, ticTacToe.Move(1, 0, 2))
	assert.EqualValues(t, 1, ticTacToe.Move(2, 1, 1))
}

func Test_TicTacToe_2(t *testing.T) {
	ticTacToe := Constructor(2)
	assert.EqualValues(t, 0, ticTacToe.Move(0, 1, 1))
	assert.EqualValues(t, 0, ticTacToe.Move(1, 1, 2))
	assert.EqualValues(t, 1, ticTacToe.Move(1, 0, 1))
}
