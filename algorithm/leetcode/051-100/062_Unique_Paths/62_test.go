package _62_Unique_Paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniquePaths(t *testing.T) {
	assert.EqualValues(t, 28, uniquePaths(3, 7))
}
