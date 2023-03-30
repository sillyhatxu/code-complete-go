package _22_Coin_Change

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_coinChange(t *testing.T) {
	assert.EqualValues(t, 20, coinChange([]int{186, 419, 83, 408}, 6249))
	assert.EqualValues(t, 2, coinChange([]int{1, 2, 5}, 6))
	assert.EqualValues(t, 3, coinChange([]int{1, 2, 5}, 11))
}
