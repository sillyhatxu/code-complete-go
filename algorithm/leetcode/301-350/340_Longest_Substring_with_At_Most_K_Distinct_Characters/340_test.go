package _40_Longest_Substring_with_At_Most_K_Distinct_Characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLongestSubstringKDistinct(t *testing.T) {
	assert.EqualValues(t, 2, lengthOfLongestSubstringKDistinct("abee", 1))
	assert.EqualValues(t, 0, lengthOfLongestSubstringKDistinct("ab", 0))
	assert.EqualValues(t, 3, lengthOfLongestSubstringKDistinct("eceba", 2))
}
