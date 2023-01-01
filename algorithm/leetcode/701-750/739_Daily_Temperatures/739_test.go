package _39_Daily_Temperatures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	//assert.EqualValues(t, []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0}, dailyTemperatures([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}))
	//assert.EqualValues(t, []int{1, 1, 4, 2, 1, 1, 0, 0}, dailyTemperatures([]int{23, 24, 25, 21, 19, 22, 26, 23}))
	//assert.EqualValues(t, []int{1, 1, 4, 2, 1, 1, 0, 0}, dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	//assert.EqualValues(t, []int{1, 1, 1, 0}, dailyTemperatures([]int{30, 40, 50, 60}))
	//assert.EqualValues(t, []int{1, 1, 0}, dailyTemperatures([]int{30, 60, 90}))
	assert.EqualValues(t, []int{2, 1, 0, 1, 0}, dailyTemperatures([]int{40, 30, 50, 20, 30}))
}
