package _31_Palindrome_Partitioning

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_partition(t *testing.T) {
	assert.EqualValues(t, [][]string{{"a", "a", "b"}, {"aa", "b"}}, partition("aab"))
}
