package _22_Best_Time_to_Buy_and_Sell_Stock_II

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	assert.EqualValues(t, 7, maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
