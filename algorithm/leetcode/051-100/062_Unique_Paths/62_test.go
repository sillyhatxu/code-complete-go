package _62_Unique_Paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniquePaths(t *testing.T) {
	assert.EqualValues(t, 28, uniquePaths(3, 7))
}

func Test_uniquePaths1(t *testing.T) {
	assert.EqualValues(t, 28, uniquePaths1(3, 7))
}

func Test_uniquePaths2(t *testing.T) {
	assert.EqualValues(t, 28, uniquePaths2(3, 7))
}
