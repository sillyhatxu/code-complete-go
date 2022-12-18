package _59_Longest_Substring_with_At_Most_Two_Distinct_Characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLongestSubstringTwoDistinct(t *testing.T) {
	assert.EqualValues(t, 3, lengthOfLongestSubstringTwoDistinct("eceba"))
	assert.EqualValues(t, 5, lengthOfLongestSubstringTwoDistinct("ccaabbb"))
}
