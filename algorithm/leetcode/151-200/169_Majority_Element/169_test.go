package _69_Majority_Element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	assert.EqualValues(t, 4, majorityElement([]int{1, 2, 3, 4, 4, 5, 6, 4, 4, 7, 4, 8, 4, 9}))
}
