package shortest_palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minInsertions(t *testing.T) {
	//TODO check
	assert.EqualValues(t, 1, minInsertions("abab"))
	assert.EqualValues(t, 2, minInsertions("abcda"))
	assert.EqualValues(t, 1, minInsertions("abab"))
	assert.EqualValues(t, 0, minInsertions("zzazz"))
	assert.EqualValues(t, 2, minInsertions("mbadm"))
	assert.EqualValues(t, 5, minInsertions("leetcode"))
}
