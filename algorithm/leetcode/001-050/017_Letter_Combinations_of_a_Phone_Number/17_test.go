package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	assert.Nil(t, letterCombinations(""))
	assert.Nil(t, letterCombinations("0"))
	assert.Nil(t, letterCombinations("1"))
	assert.EqualValues(t, []string{"a", "b", "c"}, letterCombinations("2"))
	assert.EqualValues(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
	assert.EqualValues(t, []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}, letterCombinations("234"))
}

func Test_letterCombinations1(t *testing.T) {
	assert.EqualValues(t, []string{}, letterCombinations1(""))
	assert.EqualValues(t, []string{"a", "b", "c"}, letterCombinations1("2"))
	assert.EqualValues(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations1("23"))
	assert.EqualValues(t, []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}, letterCombinations1("234"))
}

func Test_letterCombinationsOriginal(t *testing.T) {
	assert.EqualValues(t, []string{}, letterCombinationsOriginal(""))
	assert.EqualValues(t, []string{"a", "b", "c"}, letterCombinationsOriginal("2"))
	assert.EqualValues(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinationsOriginal("23"))
	assert.EqualValues(t, []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}, letterCombinationsOriginal("234"))
}
