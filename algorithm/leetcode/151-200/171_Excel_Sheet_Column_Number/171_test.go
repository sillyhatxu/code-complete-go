package _71_Excel_Sheet_Column_Number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	assert.EqualValues(t, 2147483647, titleToNumber("FXSHRXW"))
	assert.EqualValues(t, 731, titleToNumber("ABC"))
	assert.EqualValues(t, 701, titleToNumber("ZY"))
}

func Test_titleToNumber1(t *testing.T) {
	assert.EqualValues(t, 2147483647, titleToNumber1("FXSHRXW"))
	assert.EqualValues(t, 731, titleToNumber1("ABC"))
	assert.EqualValues(t, 701, titleToNumber1("ZY"))
}
