package v2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ProductOfNumbers(t *testing.T) {
	pon := Constructor()
	pon.Add(3)
	pon.Add(0)
	pon.Add(2)
	pon.Add(5)
	pon.Add(4)
	assert.EqualValues(t, 20, pon.GetProduct(2))
	assert.EqualValues(t, 40, pon.GetProduct(3))
	assert.EqualValues(t, 0, pon.GetProduct(4))
	pon.Add(8)
	assert.EqualValues(t, 32, pon.GetProduct(2))
}
