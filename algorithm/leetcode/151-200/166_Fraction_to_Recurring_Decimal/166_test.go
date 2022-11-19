package _66_Fraction_to_Recurring_Decimal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sign(t *testing.T) {
}
func Test_fractionToDecimal(t *testing.T) {
	assert.EqualValues(t, "0.(012)", fractionToDecimal(4, 333))
}
