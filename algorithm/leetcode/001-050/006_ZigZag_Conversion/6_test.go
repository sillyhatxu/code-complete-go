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
