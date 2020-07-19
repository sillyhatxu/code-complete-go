package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkCheckPermutation(b *testing.B) {
	result := CheckPermutation("dog", "god")
	assert.Equal(b, true, result)
	result = CheckPermutation("sympathetically", "optimization")
	assert.Equal(b, false, result)
	result = CheckPermutation("disease", "optimization")
	assert.Equal(b, false, result)
}
