package _79_Word_Search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
Time: O(4^1), l = len(word)
total: O(m*n*4^1)

Space: O(m*n+1)
*/
func Test_exist(t *testing.T) {
	result1 := exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED")
	assert.EqualValues(t, true, result1)
}
