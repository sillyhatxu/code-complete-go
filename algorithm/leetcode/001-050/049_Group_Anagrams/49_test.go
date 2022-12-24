package _49_Group_Anagrams

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	assert.EqualValues(t, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}, groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func Test_groupAnagramsOriginal(t *testing.T) {
	assert.EqualValues(t, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}, groupAnagramsOriginal([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
