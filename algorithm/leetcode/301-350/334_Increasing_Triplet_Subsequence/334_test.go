package _34_Increasing_Triplet_Subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_increasingTriplet(t *testing.T) {
	assert.EqualValues(t, false, increasingTriplet([]int{1, 1, -2, 6}))
}
