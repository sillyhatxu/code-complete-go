package _87_First_Unique_Character_in_a_String

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_firstUniqChar(t *testing.T) {
	assert.EqualValues(t, 0, firstUniqChar("leetcode"))
}
