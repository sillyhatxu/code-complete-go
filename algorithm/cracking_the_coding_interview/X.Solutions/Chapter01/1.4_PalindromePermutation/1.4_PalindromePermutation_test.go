package x_solutions_chapter01_PalindromePermutation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindromePermutation(t *testing.T) {
	test := map[string]bool{
		"deed":     true,
		"level":    true,
		"test":     false,
		"abac":     false,
		"aabaa":    true,
		"Tact Coa": true,
	}
	for k, v := range test {
		result := PalindromePermutation(k)
		assert.Equal(t, result, v)
	}
}
