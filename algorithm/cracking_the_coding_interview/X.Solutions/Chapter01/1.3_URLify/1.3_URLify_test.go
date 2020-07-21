package x_solutions_chapter01_URlify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestURLify2(t *testing.T) {
	input, trueLength := "Mr John Smith    ", 13
	result := URLify([]byte(input), trueLength)
	assert.Equal(t, "Mr%20John%20Smith", string(result))
}
