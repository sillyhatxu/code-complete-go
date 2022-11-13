package _34_Gas_Station

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canCompleteCircuit(t *testing.T) {
	assert.EqualValues(t, 3, canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	assert.EqualValues(t, -1, canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
