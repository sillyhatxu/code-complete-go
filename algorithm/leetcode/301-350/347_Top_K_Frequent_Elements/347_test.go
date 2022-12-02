package _47_Top_K_Frequent_Elements

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	assert.EqualValues(t, []int{1, 2}, topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
