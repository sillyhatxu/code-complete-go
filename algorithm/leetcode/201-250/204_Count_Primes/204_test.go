package _04_Count_Primes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countPrimes(t *testing.T) {
	assert.EqualValues(t, 4, countPrimes(10))
}
