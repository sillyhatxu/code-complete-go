package _16_Partition_Equal_Subset_Sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canPartition(t *testing.T) {
	assert.EqualValues(t, true, canPartition([]int{1, 5, 11, 5}))
	assert.EqualValues(t, false, canPartition([]int{1, 2, 3, 5}))
}
