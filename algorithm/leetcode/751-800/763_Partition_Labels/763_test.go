package _63_Partition_Labels

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	assert.EqualValues(t, []int{9, 7, 8}, partitionLabels("ababcbacadefegdehijhklij"))
	assert.EqualValues(t, []int{10}, partitionLabels("eccbbbbdec"))
}
