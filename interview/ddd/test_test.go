package ddd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_test(t *testing.T) {
	assert.EqualValues(t, 31, test([]int{29, -2, 4}))
}
