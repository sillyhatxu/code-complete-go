package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_convertBackups(t *testing.T) {
	assert.EqualValues(t, "0481357926", convert("0123456789", 3))
	assert.EqualValues(t, "PAYPALISHIRING", convert("PAYPALISHIRING", 1))
	assert.EqualValues(t, "ACEGBDF", convert("ABCDEFG", 2))
	assert.EqualValues(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
	assert.EqualValues(t, "AGMBFHLNCEIKOQDJP", convert("ABCDEFGHIJKLMNOPQ", 4))
}

func Test_convert1(t *testing.T) {
	assert.EqualValues(t, "0481357926", convert1("0123456789", 3))
	assert.EqualValues(t, "PAYPALISHIRING", convert1("PAYPALISHIRING", 1))
	assert.EqualValues(t, "ACEGBDF", convert1("ABCDEFG", 2))
	assert.EqualValues(t, "PAHNAPLSIIGYIR", convert1("PAYPALISHIRING", 3))
	assert.EqualValues(t, "AGMBFHLNCEIKOQDJP", convert1("ABCDEFGHIJKLMNOPQ", 4))
}
