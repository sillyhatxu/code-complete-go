package _21_Best_Time_to_Buy_and_Sell_Stock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	assert.EqualValues(t, 17, maxProfit([]int{7, 3, 20, 3, 1, 2}))
	assert.EqualValues(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
